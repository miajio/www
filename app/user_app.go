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
		email := ctx.DefaultPostForm("email", "")
		password := ctx.DefaultPostForm("password", "")
		if email == "" {
			ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "邮箱不能为空", "error": "用户名不能为空"})
			return
		}
		if password == "" {
			ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "密码不能为空", "error": "密码不能为空"})
			return
		}

		searchSQL := "SELECT `uid`, `username`, `head_pic` FROM `user_info` WHERE `email` = ? and `password` = MD5(?) and status = 1"

		type obj struct {
			Uid      string `db:"uid" json:"uid"`
			UserName string `db:"username" json:"username"`
			HeadPic  string `db:"head_pic" json:"headPic"`
		}

		var result obj

		if err := lib.DB.Get(&result, searchSQL, email, password); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "用户名或密码错误", "error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": result})
		// ctx.HTML(http.StatusOK, "r9.html", gin.H{"uid": uid})
	})

}
