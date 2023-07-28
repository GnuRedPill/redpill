package dao

import (
	"app/models"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
)

type RLogDao struct {
	DB *bun.DB
}

func (a RLogDao) CreatRLog(context context.Context, url, content string) (models.RLog, error) {
	rlog := new(models.RLog)
	rlog.CreateTime = time.Now()
	rlog.UpdateTime = time.Now()
	rlog.State = models.State_Raw
	rlog.ReadUrl = url
	rlog.Content = content
	_, err := a.DB.NewInsert().Model(rlog).Exec(context, rlog)
	fmt.Println(rlog.ID, rlog.ReadUrl)
	return *rlog, err
}

func (a RLogDao) Done(context context.Context, id int) {
	rlog := new(models.RLog)
	rlog.State = models.State_Parsed
	a.DB.NewUpdate().Model(rlog).Column("state").Where("id = ?", id).Exec(context)
}

func (a RLogDao) GetLastRLog(context context.Context) (models.RLog, error) {
	rlog := new(models.RLog)
	_, err := a.DB.NewSelect().Model(rlog).Where("state =  ?", models.State_Raw).Order("id desc").Exec(context, rlog)
	return *rlog, err
}

func (a RLogDao) GetLastRLogCount(context context.Context) int {
	rlog := new(models.RLog)
	count, err := a.DB.NewSelect().Model(rlog).Where("state =  ?", models.State_Raw).Count(context)
	if err != nil {
		log.Println(err)
		return 0
	}
	return count
}
