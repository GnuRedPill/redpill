package main

import (
	"app/action"
	"app/db"
	"app/utils"
	"net/http"
)

func main() {
	utils.InitLog()
	db := db.InitMysql()
	defer db.Close()
	router := action.InitRoute(db)
	http.ListenAndServe("0.0.0.0:8080", router)
}
