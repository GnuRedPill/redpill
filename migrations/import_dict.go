package main

import (
	"app/db"
	"app/models"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func ImportDict() {
	log.SetFlags(log.Ldate & log.Lshortfile)
	ImportPositively()
	ImportNegative()
	ImportSenticent()
}

func ImportSenticent() {
	db := db.InitMysql()
	defer db.Close()
	f, err := os.Open("data/senticnet.txt")
	defer f.Close()
	if err != nil {
		log.Println(err)
	}
	c := csv.NewReader(f)
	arr, err := c.ReadAll()
	if err != nil {
		log.Println(err)
	}
	are := []models.Tag{}
	for i := range arr {
		item := arr[i]
		fmt.Println(item[0], item[1])
		tag := models.Tag{}
		if item[0] == "positive" {
			tag.Type = models.TYPE_Positively
		}
		if item[0] == "negative" {
			tag.Type = models.TYPE_Negative
		}
		tag.CreateTime = time.Now()
		tag.Name = item[1]
		tag.State = models.State_Raw
		tag.Emo = ""
		are = append(are, tag)
	}
	db.NewInsert().Model(&are).Exec(context.Background())
}

func ImportPositively() {
	db := db.InitMysql()
	defer db.Close()
	arr := openCsv(`data\positive.txt`)
	tags := []models.Tag{}
	for i := range arr {
		v := arr[i]
		fmt.Println(v)
		tag := models.Tag{}
		tag.Type = models.TYPE_Positively
		tag.CreateTime = time.Now()
		tag.Name = v
		tag.State = models.State_Raw
		tag.Emo = ""
		tags = append(tags, tag)
	}
	db.NewInsert().Model(&tags).Exec(context.Background())
}
func ImportNegative() {
	db := db.InitMysql()
	defer db.Close()
	arr := openCsv(`data\negative.txt`)
	tags := []models.Tag{}
	for i := range arr {
		v := arr[i]
		fmt.Println(v)
		tag := models.Tag{}
		tag.Type = models.TYPE_Negative
		tag.CreateTime = time.Now()
		tag.Name = v
		tag.State = models.State_Raw
		tag.Emo = ""
		tags = append(tags, tag)
	}
	db.NewInsert().Model(&tags).Exec(context.Background())
}
func openCsv(name string) []string {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}
	brr, err := io.ReadAll(f)
	if err != nil {
		log.Println(err)
	}
	str := string(brr)
	arr := strings.Split(str, "\n")
	return arr
}
