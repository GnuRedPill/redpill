package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

const (
	TYPE_Positively = 1 //正面
	TYPE_Negative   = 2 //负面
	TYPE_RAW        = 3

	EMO_apprehensive = "apprehensive" //焦虑
	EMO_sorrowful    = "sorrowful"    //悲伤
	EMO_assault      = "assault"      //人身攻击

	State_Raw    = 1 //原始分词
	State_Parsed = 2 //已解析
)

type Tag struct {
	bun.BaseModel `bun:"table:tags"`
	ID            int    `bun:",pk,autoincrement"`
	Name          string //tag 名字
	Type          int    //情绪趋向
	Emo           string //具体情绪
	State         int
	CreateTime    time.Time
}

type L2T struct {
	bun.BaseModel `bun:"table:l2tags"`
	ID            int `bun:",pk,autoincrement"`
	Tag           int //标签id
	Log           int //日志
}

var _ bun.BeforeAppendModelHook = (*Tag)(nil)

func (m *Tag) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreateTime = time.Now()
	case *bun.UpdateQuery:

	}
	return nil
}
