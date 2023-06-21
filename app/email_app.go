package app

import (
	"github.com/gin-gonic/gin"
	"github.com/miajio/www/app/logic"
	"github.com/miajio/www/bin"
)

type emailRouterImpl struct{}

type emailRouter bin.Router

// EmailRouter 邮件路由
var EmailRouter emailRouter = (*emailRouterImpl)(nil)

func (*emailRouterImpl) Running(e *gin.Engine) {
	emailGroup := e.Group("/email")
	emailGroup.POST("/sendCheckCode", logic.EmailLogic.SendCheckCode)
}
