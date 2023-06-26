package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miajio/www/lib"
	"github.com/miajio/www/req"
)

type articleLogicImpl struct{}

type articleLogic interface {
	Find(ctx *gin.Context) // Find 获取文章列表
}

var ArticleLogic articleLogic = (*articleLogicImpl)(nil)

// Find 获取文章列表
func (*articleLogicImpl) Find(ctx *gin.Context) {
	request := req.ArticleFindRequest{}
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误", "error": lib.TransError(err)})
		return
	}
	request.Default()

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": request.Group})
}
