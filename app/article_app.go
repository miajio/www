package app

import (
	"github.com/gin-gonic/gin"
	"github.com/miajio/www/app/logic"
	"github.com/miajio/www/bin"
)

type articleRouterImpl struct{}

type articleRouter bin.Router

var ArticleRouter articleRouter = (*articleRouterImpl)(nil)

func (*articleRouterImpl) Running(e *gin.Engine) {
	articleGroup := e.Group("/article")
	articleGroup.GET("/find", logic.ArticleLogic.Find)
}
