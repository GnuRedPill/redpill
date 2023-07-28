package main

import (
	"app/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestParseSen(t *testing.T) {

	f, err := os.Open("../data/senticnet.txt")
	defer f.Close()
	if err != nil {
		log.Println(err)
	}
	c := csv.NewReader(f)
	arr, err := c.ReadAll()
	if err != nil {
		log.Println(err)
	}
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
		//db.NewInsert().Model(&tag).Exec(context.Background())
	}

}
