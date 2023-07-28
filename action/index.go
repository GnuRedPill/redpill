package action

import (
	"app/dao"
	"app/utils"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bunrouter"
)

type IndexAction struct {
	db *bun.DB
}

func (i *IndexAction) Index(w http.ResponseWriter, req bunrouter.Request) error {
	w.Write([]byte("go cgi 0.01"))
	return nil
}

type PushText struct {
	ReadUrl, Content string
}

func (i *IndexAction) PushText(w http.ResponseWriter, req bunrouter.Request) error {
	fmt.Println("is  req ====>")
	token := req.URL.Query().Get("token")
	res := utils.ApiResult{}
	if token != "xxxxx" {
		w.Write(res.Fail().Json())
		return nil
	}
	ud := dao.RLogDao{DB: i.db}
	p := PushText{}
	brr, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(brr, &p)
	if err != nil {
		log.Println(err)
	}
	ud.CreatRLog(context.Background(), p.ReadUrl, p.Content)
	w.Write(res.Success().Json())
	return nil
}

func (i *IndexAction) UploadImage(w http.ResponseWriter, req bunrouter.Request) error {
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Println(err)
	}
	allimage := req.MultipartForm.File["image"][0]
	f, err := allimage.Open()
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	brr, err := io.ReadAll(f)
	if err != nil {
		log.Println(err)
	}
	res := utils.ApiResult{}
	ok, types := utils.ImageUtils{}.GetImageType(brr)
	if ok {
		name := utils.ImageUtils{}.ImageToMd5(brr) + "." + types
		fd, err := dao.FileDao{DB: i.db}.CreatFile(req.Context(), name)
		if err != nil {
			log.Println(err)
		}
		fdf, err := os.Create("files/" + fd.Ident)
		defer fdf.Close()
		if err != nil {
			log.Println(err)
		}
		fdf.Write(brr)
		res.SuccessAData(fd)

	}

	w.Write(res.Json())
	return nil
}
