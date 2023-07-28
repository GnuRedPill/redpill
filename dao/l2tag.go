package dao

import (
	"app/models"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/uptrace/bun"
)

type R2tagDao struct {
	DB *bun.DB
}

func (r *R2tagDao) CreateTag(tag string) (models.Tag, error) {
	tagitem := models.Tag{}
	tagitem.CreateTime = time.Now()
	tagitem.State = models.State_Raw
	tagitem.Type = models.TYPE_RAW
	tagitem.Name = tag
	r.DB.NewInsert().Model(&tagitem).Exec(context.Background())
	return tagitem, nil
}

func (r *R2tagDao) GetTag(tags string) (models.Tag, error) {
	tag := models.Tag{}
	_, err := r.DB.NewSelect().Model(tag).Where("name = ?", tag).Exec(context.Background())
	return tag, err
}

func (r *R2tagDao) GetTags(tags []string) []models.Tag {
	arr := []models.Tag{}
	for i := range tags {
		tagitem, err := r.GetTag(tags[i])
		if err != nil {
			if err == sql.ErrNoRows {
				tagitem, _ = r.CreateTag(tags[i])
				arr = append(arr, tagitem)
			} else {
				log.Println(err)
			}
		} else {
			if tagitem.ID > 0 {
				arr = append(arr, tagitem)
			} else {

				tagitem, _ = r.CreateTag(tags[i])
				arr = append(arr, tagitem)

			}
		}
	}

	return arr
}

func (r *R2tagDao) SaveL2T(logid int, tags []models.Tag) {
	for i := range tags {
		t := tags[i]
		if t.ID > 0 {

			l2titem := models.L2T{}
			l2titem.Tag = t.ID
			l2titem.Log = logid
			r.DB.NewInsert().Model(&l2titem).Exec(context.Background())

		}

	}
}
