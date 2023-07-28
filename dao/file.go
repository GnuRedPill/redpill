package dao

import (
	"app/models"
	"context"
	"time"

	"github.com/uptrace/bun"
)

type FileDao struct {
	DB *bun.DB
}

func (a FileDao) CreatFile(context context.Context, Ident string) (models.Files, error) {
	f := new(models.Files)
	f.Ident = Ident
	f.CreateTime = time.Now()
	f.UpdateTime = time.Now()
	_, err := a.DB.NewInsert().Model(f).Exec(context, f)
	return *f, err
}
