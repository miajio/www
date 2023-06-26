package logic

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miajio/www/lib"
	"github.com/miajio/www/req"
	"github.com/miajio/www/store"
)

type userLogicImpl struct{}

type userLogic interface {
	Login(*gin.Context)    // Login 登录
	Register(*gin.Context) // Register 注册
	Auth(*gin.Context)     // Auth Token鉴权
}

// UserLogic 用户业务逻辑
var UserLogic userLogic = (*userLogicImpl)(nil)

// Login 登录
func (*userLogicImpl) Login(ctx *gin.Context) {
	var request req.UserEmailLoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误", "error": lib.TransError(err)})
		return
	}

	result, err := store.UserInfoStore.Login(request.Email, request.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "用户名或密码错误", "error": err.Error()})
		return
	}
	msg, err := lib.Jwt.GenerateToken([]byte(lib.ServerCfg.JwtKey), result, 24*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "jwt生成异常", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": msg})
}

// Register 注册
func (*userLogicImpl) Register(ctx *gin.Context) {
	var request req.UserEmailRegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误", "error": lib.TransError(err)})
		return
	}

	if err := EmailLogic.CheckCode(request.Email, "register", request.Uid, request.CheckCode); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "验证码错误", "error": err.Error()})
		return
	}

	ok, err := store.UserInfoStore.EmailRepeat(request.Email)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "服务器异常", "error": "服务器异常:" + err.Error()})
		return
	}
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "当前邮箱已被注册", "error": "当前邮箱已被注册"})
		return
	}

	result, err := store.UserInfoStore.Register(request.Username, request.Email, request.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "注册失败", "error": err.Error()})
		return
	}
	msg, err := lib.Jwt.GenerateToken([]byte(lib.ServerCfg.JwtKey), result, 24*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "jwt生成异常", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功", "data": msg})
}

// Auth Token 鉴权
func (*userLogicImpl) Auth(ctx *gin.Context) {
	tk := ctx.Request.Header.Get("Authorization")
	msg, err := lib.Jwt.ParseToken([]byte(lib.ServerCfg.JwtKey), tk)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "jwt鉴权失败,请重新登录", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "鉴权成功", "data": msg})
}
