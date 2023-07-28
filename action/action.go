package action

import (
	"app/utils"

	"github.com/uptrace/bun"
	"github.com/uptrace/bunrouter"
)

func InitRoute(db *bun.DB) *bunrouter.Router {
	router := bunrouter.New(
		bunrouter.Use(utils.PrintMiddleware),
	)
	indexAction := IndexAction{db: db}
	router.GET("/", indexAction.Index)
	router.POST("/upload", indexAction.UploadImage)
	router.POST("/push_text", indexAction.PushText)
	return router
}
