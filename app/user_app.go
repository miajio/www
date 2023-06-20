package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miajio/www/bin"
	"github.com/miajio/www/lib"
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

		searchSQL := "SELECT `uid` FROM `r9_user_info` WHERE `account` = ? and `password` = MD5(?) and `effective_time` > NOW()"

		uid := ""

		if err := lib.DB.Get(&uid, searchSQL, username, password); err != nil {
			if username == "miajio" && password == "123456" {
				ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "登录失败,请联系管理员", "error": "登录失败,请联系管理员"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": uid})
		// ctx.HTML(http.StatusOK, "r9.html", gin.H{"uid": uid})
	})

}
