package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miajio/www/bin"
)

type userRouterImpl struct{}

type userRouter bin.Router

// UserRouter 用户路由
var UserRouter userRouter = (*userRouterImpl)(nil)

func (*userRouterImpl) Running(e *gin.Engine) {
	userGroup := e.Group("/user")
	userGroup.POST("/login", func(ctx *gin.Context) {
		username := ctx.DefaultPostForm("username", "")
		password := ctx.DefaultPostForm("password", "")
		if username == "" {
			ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "用户名不能为空", "error": "用户名不能为空"})
			return
		}
		if password == "" {
			ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "密码不能为空", "error": "密码不能为空"})
			return
		}

		if username == "miajio" && password == "123456" {
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "登录失败,用户名或密码错误", "error": "登录失败,用户名或密码错误"})
		return
	})
}
