package utils

import (
	"encoding/json"
	"log"
)

const (
	API_SUCCESS = 1
	API_FAIL    = 2
)

type ApiResult struct {
	Code    int
	Message string
	Data    interface{}
}

func (a *ApiResult) Success() ApiResult {
	a.Code = API_SUCCESS
	a.Message = "success"
	return *a
}
func (a *ApiResult) SuccessAData(data interface{}) ApiResult {
	a.Code = API_SUCCESS
	a.Message = "success"
	a.Data = data
	return *a
}

func (a *ApiResult) SuccessAndMessage(msg string) ApiResult {
	a.Code = API_SUCCESS
	a.Message = msg
	return *a
}

func (a *ApiResult) SuccessAmsgADate(msg string, data interface{}) ApiResult {
	a.Code = API_SUCCESS
	a.Message = msg
	a.Data = data
	return *a
}

func (a *ApiResult) Fail() ApiResult {
	a.Code = API_FAIL
	a.Message = "fail"
	return *a
}
func (a *ApiResult) FailAData(data interface{}) ApiResult {
	a.Code = API_FAIL
	a.Message = "fail"
	a.Data = data
	return *a
}

func (a *ApiResult) FailAndMessage(msg string) ApiResult {
	a.Code = API_FAIL
	a.Message = msg
	return *a
}

func (a *ApiResult) FailAndMessageAData(msg string, data interface{}) ApiResult {
	a.Code = API_FAIL
	a.Message = msg
	a.Data = data
	return *a

}
func (a ApiResult) Json() []byte {
	brr, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
	}
	return brr
}
