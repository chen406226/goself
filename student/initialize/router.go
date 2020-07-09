package initialize

import (
	"github.com/gin-gonic/gin"
	"student/global"
	"student/router"

	//"student/router"
)

func Router() *gin.Engine {
	Router := gin.Default()
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GL_LOG.Debug("use middleware logger")
	// 跨域
	//Router.Use(middleware.Cors())
	//global.GVA_LOG.Debug("use middleware cors")
	ApiGroup := Router.Group("")
	//router.InitUserRouter(ApiGroup)							//注册用户路由
	router.InitBaseRouter(ApiGroup)							// 注册基础功能路由 不做鉴权
	global.GL_LOG.Info("router register success")
	return Router
}