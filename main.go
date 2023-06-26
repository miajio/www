package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/miajio/www/app"
	"github.com/miajio/www/bin"
	"github.com/miajio/www/lib"
	"github.com/miajio/www/log"
	"github.com/spf13/viper"
)

type (
	// errorRequest 错误消息请求体
	errorRequest struct {
		Code string `json:"code" form:"code" uri:"code" binding:"required"`
		Msg  string `json:"msg" form:"msg" uri:"msg" binding:"required"`
		Err  string `json:"err" form:"err" uri:"err"`
	}
)

var (
	// GinRouters gin路由集
	GinRouters bin.Routers = []bin.Router{
		app.UserRouter,    // 用户路由
		app.MsgRouter,     // 消息路由
		app.EmailRouter,   // 邮件路由
		app.ArticleRouter, // 文章路由
	}

	errorHandler = func(ctx *gin.Context) {
		req := errorRequest{}

		code := 500
		errMsg := ""
		if err := ctx.ShouldBind(&req); err != nil {
			errMsg = lib.TransError(err)
		} else {
			errMsg = strings.Join([]string{req.Msg, req.Err}, " <br /> ^1000")
		}

		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"title": code,
			"err":   errMsg,
		})
	}
)

func init() {
	lp := log.LoggerParam{
		Path:       "./log",
		MaxSize:    256,
		MaxBackups: 10,
		MaxAge:     7,
		Compress:   false,
	}
	lib.Log = lp.Default()

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	runningPath, _ := os.Getwd()

	v.AddConfigPath(runningPath)

	if err := v.ReadInConfig(); err != nil {
		lib.Log.Errorf("读取配置文件失败: %v", err)
		os.Exit(0)
	}

	if err := v.UnmarshalKey("mysql", &lib.DBCfg); err != nil {
		lib.Log.Errorf("数据库配置读取失败: %v", err)
		os.Exit(0)
	} else {
		client, err := lib.MysqlUtil.Connect(lib.DBCfg)
		if err != nil {
			lib.Log.Errorf("数据库连接失败: %v", err)
			os.Exit(0)
		}
		lib.DB = client
	}

	// Redis服务
	if err := v.UnmarshalKey("redis", &lib.RedisCfg); err != nil {
		lib.Log.Errorf("redis配置读取失败: %v", err)
		os.Exit(0)
	} else {
		// 获取redis client
		lib.RedisClient = lib.RedisCfg.NewClient()
		val, err := lib.RedisClient.Ping().Result()
		if err != nil {
			lib.Log.Errorf("redis ping fail: %v", err)
			os.Exit(0)
		}
		fmt.Println(val)
	}

	// 邮箱服务
	if err := v.UnmarshalKey("email", &lib.EmailCfg); err != nil {
		lib.Log.Errorf("邮件配置读取失败: %v", err)
	} else {
		lib.EmailCfg.IsStatus = true
	}

	if err := v.UnmarshalKey("server", &lib.ServerCfg); err != nil {
		lib.Log.Errorf("系统服务配置读取失败: %v", err)
		os.Exit(0)
	} else {
		if lib.ServerCfg.Port[0] != ':' {
			lib.ServerCfg.Port = ":" + lib.ServerCfg.Port
		}

		if strings.TrimSpace(lib.ServerCfg.HttpsKey) == "" || strings.TrimSpace(lib.ServerCfg.HttpsPem) == "" {
			lib.ServerCfg.UseHttps = false
		}
	}

	lib.Trans = lib.InitValidateTrans(binding.Validator.Engine().(*validator.Validate))
}

func main() {
	e := gin.New()

	e.Static("/static/images", "./static/images")
	// e.Static("/static/images/head", "./static/images/head")

	e.Static("/static/js", "./static/js")
	e.Static("/static/css", "./static/css")

	e.Static("/static/bootstrap/js", "./static/bootstrap/js")
	e.Static("/static/bootstrap/css", "./static/bootstrap/css")

	bin.GinUtil.LoadHTMLFolders(e, []string{"./static"}, ".html")
	bin.GinUtil.LoadRouters(e, GinRouters...)

	e.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// error
	e.GET("/error", errorHandler)
	e.POST("/error", errorHandler)

	e.GET("/:page", func(ctx *gin.Context) {
		page := ctx.Param("page")
		if strings.Contains(page, ".html") {
			ctx.HTML(http.StatusOK, page, nil)
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "page is not found",
		})
	})

	if lib.ServerCfg.UseHttps {
		go func() {
			e1 := gin.New()
			e1.GET("/", func(ctx *gin.Context) {
				ctx.Redirect(http.StatusMovedPermanently, lib.ServerCfg.HttpsHost)
			})

			e1.GET("/:page", func(ctx *gin.Context) {
				ctx.Redirect(http.StatusMovedPermanently, lib.ServerCfg.HttpsHost)
			})
			e1.Use(bin.Limit(64))

			if err := e1.Run(":80"); err != nil {
				lib.Log.Errorf("init gin 80 port server fail: %v", err)
			}
		}()
		e.Use(bin.Limit(64))
		e.RunTLS(lib.ServerCfg.Port, lib.ServerCfg.HttpsPem, lib.ServerCfg.HttpsKey)
		return
	}
	e.Use(bin.Limit(64))
	e.Run(lib.ServerCfg.Port)
}
