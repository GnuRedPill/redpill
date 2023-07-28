package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"strings"
)

type ImageUtils struct{}

var (
	ImageTypes = []string{
		"png",
		"jpeg",
		"gif",
	}
)

// buff make([]byte,512)
func (ImageUtils) GetImageType(buff []byte) (bool, string) {
	filetype := http.DetectContentType(buff)
	arr := strings.Split(filetype, "/")
	if arr[0] == "image" && arrin(ImageTypes, arr[1]) {
		return true, arr[1]
	} else {
		return false, arr[1]
	}
}

func (ImageUtils) ImageToBase64(buff []byte) string {
	return base64.StdEncoding.EncodeToString(buff)
}

func (ImageUtils) ImageToMd5(buff []byte) string {
	md5arr := md5.Sum(buff)
	md5st := md5arr[0:16]
	return hex.EncodeToString(md5st)
}

func arrin(arrs []string, item string) bool {
	for _, v := range arrs {
		if strings.EqualFold(v, item) {
			return true
		}
	}
	return false
}
