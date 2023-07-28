package main

import (
	"app/dao"
	"app/db"
	"app/utils"
	"app/utils/cut"
	"context"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type CutUtilsTask struct {
	CutUtils *cut.Segmenter
}

func (c *CutUtilsTask) Init() {
	c.CutUtils = cut.NewSegmenter()
	f, err := os.Open("data/cut_dict.txt")
	if err != nil {
		panic(err)
	}
	brr, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(brr), "\n")
	garr := []cut.GseToken{}
	for i := range arr {
		token := arr[i]
		if len(token) > 0 {
			garr = append(garr, cut.GseToken{
				Text: token,
				Freq: 2,
				Pos:  cut.NOUN,
			},
			)
		}

	}
	c.CutUtils.AddToken(garr)
}

func main() {
	utils.InitLog()
	db := db.InitMysql()
	defer db.Close()
	ct := CutUtilsTask{}
	ct.Init()
	rlogdao := dao.RLogDao{DB: db}
	l2tagdao := dao.R2tagDao{DB: db}
	for {
		if rlogdao.GetLastRLogCount(context.Background()) > 0 {

			rlog, err := rlogdao.GetLastRLog(context.Background())
			if err != nil {
				log.Println(err)
				time.Sleep(5 * time.Second)
			}
			cutarrs := ct.CutUtils.Cut(rlog.Content)
			//fmt.Println("====> cut dict", cutarrs)
			tags := l2tagdao.GetTags(cutarrs)
			//fmt.Println("=====> get tag obj or  create", tags)
			l2tagdao.SaveL2T(rlog.ID, tags)

			rlogdao.Done(context.Background(), rlog.ID)

		} else {
			time.Sleep(10 * time.Second)
		}
	}
}
