package app

import (
	"github.com/gin-gonic/gin"
	"github.com/miajio/www/app/logic"
	"github.com/miajio/www/bin"
)

type userRouterImpl struct{}

type userRouter bin.Router

// UserRouter 用户路由
var UserRouter userRouter = (*userRouterImpl)(nil)

func (*userRouterImpl) Running(e *gin.Engine) {
	userGroup := e.Group("/user")
	userGroup.POST("/login", logic.UserLogic.Login)
	userGroup.POST("/register", logic.UserLogic.Register)

}
