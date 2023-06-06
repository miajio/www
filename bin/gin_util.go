package bin

import "github.com/gin-gonic/gin"

// Router 基于gin的路由抽象
type Router interface {
	Running(*gin.Engine)
}

// Routers 基于gin的路由集
type Routers []Router

type ginUtilImpl struct{}

type ginUtil interface {
	LoadHTMLFolders(e *gin.Engine, folders []string, suffix string) // LoadHTMLFolders 装载html文件
	LoadRouters(e *gin.Engine, routers ...Router)                   // LoadRouters 装载路由集
}

// GinUtil gin工具实现
var GinUtil ginUtil = (*ginUtilImpl)(nil)

// LoadHTMLFolders 装载html文件
func (*ginUtilImpl) LoadHTMLFolders(e *gin.Engine, folders []string, suffix string) {
	if folders == nil {
		return
	}

	files := make([]string, 0)

	for _, folder := range folders {
		inFiles, err := FileUtil.FindFileByFolderChildren(folder, suffix)
		if err != nil {
			continue
		}
		files = append(files, inFiles...)
	}

	files = SliceDistinct(files...)
	if files != nil {
		e.LoadHTMLFiles(files...)
	}
}

// LoadRouters 装载路由集
func (*ginUtilImpl) LoadRouters(e *gin.Engine, routers ...Router) {
	for i := range routers {
		routers[i].Running(e)
	}
}
