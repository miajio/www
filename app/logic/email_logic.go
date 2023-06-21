package logic

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miajio/www/lib"
	"github.com/miajio/www/req"
	"github.com/miajio/www/store"
)

type emailLogicImpl struct{}

type emailLogic interface {
	SendCheckCode(ctx *gin.Context)                     // SendCheckCode 发送验证码
	CheckCode(email, emailType, uid, code string) error // CheckCode 校验验证码
	DelCheckCode(email, emailType, uid string) error    // DelCheckCode 删除当前验证码
}

// EmailLogic 邮箱业务
var EmailLogic emailLogic = (*emailLogicImpl)(nil)

// SendCheckCode 发送验证码
func (e *emailLogicImpl) SendCheckCode(ctx *gin.Context) {
	request := req.EmailSendCheckCodeRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误", "error": lib.TransError(err)})
		return
	}

	switch request.EmailType {
	case "register":
		ok, err := store.UserInfoStore.EmailRepeat(request.Email)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "服务器异常", "error": "服务器异常:" + err.Error()})
			return
		}
		if ok {
			ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "当前邮箱已被注册", "error": "当前邮箱已被注册"})
			return
		}
		uid, err := e.register(request.Email)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "服务器异常", "error": "服务器异常:" + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "发送成功", "data": uid})
	case "login":
		ok, err := store.UserInfoStore.EmailRepeat(request.Email)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "服务器异常", "error": "服务器异常:" + err.Error()})
			return
		}
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "当前邮箱未注册", "error": "当前邮箱未注册"})
			return
		}
		uid, err := e.login(request.Email)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "服务器异常", "error": "服务器异常:" + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "发送成功", "data": uid})
	case "update":
		ok, err := store.UserInfoStore.EmailRepeat(request.Email)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "服务器异常", "error": "服务器异常:" + err.Error()})
			return
		}
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "当前邮箱未注册", "error": "当前邮箱未注册"})
			return
		}
		uid, err := e.update(request.Email)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": "服务器异常", "error": "服务器异常:" + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "发送成功", "data": uid})
	default:
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误", "error": "未知的邮件推送请求"})
	}
}

// CheckCode 校验验证码
func (e *emailLogicImpl) CheckCode(email, emailType, uid, code string) error {
	key, err := e.getKey(emailType, email, uid)
	if err != nil {
		return err
	}
	rcode, err := lib.RedisClient.Get(key).Result()
	if err != nil {
		return err
	}
	if rcode != code {
		return errors.New("验证码错误")
	}
	return nil
}

// DelCheckCode 删除当前验证码
func (e *emailLogicImpl) DelCheckCode(email, emailType, uid string) error {
	key, err := e.getKey(emailType, email, uid)
	if err != nil {
		return err
	}
	_, err = lib.RedisClient.Del(key).Result()
	return err
}

// register 注册邮件发送逻辑
func (e *emailLogicImpl) register(email string) (string, error) {
	var msg strings.Builder
	msg.WriteString("miajio.com:\n")
	msg.WriteString("    您好!感谢您注册miajio平台,您的验证码是: %s\n")
	msg.WriteString("    当前验证码的有效时间是: 30分钟\n")
	msg.WriteString("如果非您本人操作,请删除该邮件并不要将验证码告诉给他人!")
	m := msg.String()
	code := lib.RandCheckCode(6)
	m = fmt.Sprintf(m, code)

	uid := lib.UID()

	if err := lib.EmailCfg.Send(email, "欢迎您注册miajio平台", m); err != nil {
		return "", err
	}
	key, _ := e.getKey("register", email, uid)
	if err := lib.RedisClient.SetNX(key, code, time.Minute*30).Err(); err != nil {
		return "", err
	}
	return uid, nil
}

func (e *emailLogicImpl) login(email string) (string, error) {
	var msg strings.Builder
	msg.WriteString("miajio.com:\n")
	msg.WriteString("    您好!您当前正在使用邮箱验证码登录,您的验证码是: %s\n")
	msg.WriteString("    当前验证码的有效时间是: 30分钟\n")
	msg.WriteString("如果非您本人操作,请删除该邮件并不要将验证码告诉给他人!")
	m := msg.String()
	code := lib.RandCheckCode(6)
	m = fmt.Sprintf(m, code)

	uid := lib.UID()

	if err := lib.EmailCfg.Send(email, "欢迎您使用miajio平台", m); err != nil {
		return "", err
	}
	key, _ := e.getKey("login", email, uid)
	if err := lib.RedisClient.SetNX(key, code, time.Minute*30).Err(); err != nil {
		return "", err
	}
	return uid, nil
}

func (e *emailLogicImpl) update(email string) (string, error) {
	var msg strings.Builder
	msg.WriteString("miajio.com:\n")
	msg.WriteString("    您好!您当前正在申请修改关键信息或密码,您的验证码是: %s\n")
	msg.WriteString("    当前验证码的有效时间是: 30分钟\n")
	msg.WriteString("如果非您本人操作,请删除该邮件并不要将验证码告诉给他人!")
	m := msg.String()
	code := lib.RandCheckCode(6)
	m = fmt.Sprintf(m, code)

	uid := lib.UID()

	if err := lib.EmailCfg.Send(email, "欢迎您使用miajio平台", m); err != nil {
		return "", err
	}

	key, _ := e.getKey("update", email, uid)
	if err := lib.RedisClient.SetNX(key, code, time.Minute*30).Err(); err != nil {
		return "", err
	}
	return uid, nil
}

func (*emailLogicImpl) getKey(emailType, email, uid string) (key string, err error) {
	switch emailType {
	case "register":
		key = "EMAIL:REGISTER:" + email + ":" + uid
	case "login":
		key = "EMAIL:LOGIN:" + email + ":" + uid
	case "update":
		key = "EMAIL:UPDATE:" + email + ":" + uid
	default:
		err = errors.New("未知的校验方式")
	}
	return
}
