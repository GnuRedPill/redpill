package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type RLog struct {
	bun.BaseModel `bun:"table:rlogs"`
	ID            int `bun:",pk,autoincrement"`
	ReadUrl       string
	Content       string
	State         int
	CreateTime    time.Time
	UpdateTime    time.Time
}

var _ bun.BeforeAppendModelHook = (*RLog)(nil)

func (m *RLog) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreateTime = time.Now()
	case *bun.UpdateQuery:
		m.UpdateTime = time.Now()
	}
	return nil
}
