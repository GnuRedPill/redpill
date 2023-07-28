package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Files struct {
	bun.BaseModel `bun:"table:files"`
	ID            int `bun:",pk,autoincrement"`
	Ident         string
	CreateTime    time.Time
	UpdateTime    time.Time
}

var _ bun.BeforeAppendModelHook = (*Files)(nil)

func (m *Files) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreateTime = time.Now()
	case *bun.UpdateQuery:
		m.UpdateTime = time.Now()
	}
	return nil
}
