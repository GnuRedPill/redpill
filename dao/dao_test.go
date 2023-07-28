package dao

import (
	"app/db"
	"app/utils"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFileUpload(t *testing.T) {
	db := db.InitMysql()
	res := utils.ApiResult{}
	defer db.Close()
	ud := FileDao{DB: db}
	f, err := os.Open("dao.go")
	defer f.Close()
	brr, err := io.ReadAll(f)
	bbrr := base64.StdEncoding.EncodeToString(brr)
	fmt.Println(bbrr)
	md5arr := md5.Sum(brr)
	md5st := md5arr[0:16]
	fmt.Println(md5st)
	u, err := ud.CreatFile(context.Background(), hex.EncodeToString(md5st))
	if err != nil {
		t.Log(err)
	}
	t.Log(string(res.SuccessAData(u).Json()))
}

func TestAlayImage(t *testing.T) {
	imagepath := `C:\Users\Administrator\Downloads\Keith.jpg`
	f, _ := os.Open(imagepath)
	buff := make([]byte, 512)
	f.Read(buff)
	ok, typee := utils.ImageUtils{}.GetImageType(buff)
	fmt.Println(ok, typee)
}

func TestCreatFileLocal(t *testing.T) {
	imagepath := `C:\Users\Administrator\Downloads\Keith.jpg`
	f, _ := os.Open(imagepath)
	defer f.Close()
	brr, _ := io.ReadAll(f)

	f2, err := os.Create("../files/" + "xxxx" + "." + "jpeg")
	defer f2.Close()
	if err != nil {
		panic(err)
	}
	f2.Write(brr)
}

func TestRlog(t *testing.T) {
	res := utils.ApiResult{}
	db := db.InitMysql()
	defer db.Close()
	ud := RLogDao{DB: db}
	u, _ := ud.CreatRLog(context.Background(), "http://baidu.com", "aaaa")
	u, _ = ud.GetLastRLog(context.Background())
	t.Log(string(res.SuccessAData(u).Json()))
}
