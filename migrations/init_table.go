package main

import (
	"app/db"
	"app/models"
	"context"
	"fmt"
)

func InitTable() {
	fmt.Println("init table")
	db := db.InitMysql()
	defer db.Close()
	// _, err := db.NewCreateTable().Model((*models.Tag)(nil)).IfNotExists().Exec(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	_, err := db.NewCreateTable().Model((*models.L2T)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	// _, err = db.NewCreateTable().Model((*models.RLog)(nil)).IfNotExists().Exec(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err = db.NewCreateTable().Model((*models.Files)(nil)).IfNotExists().Exec(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
