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
	router.InitBaseRouter(ApiGroup)							// 注册基础功能路由 不做鉴权
	router.InitUserRouter(ApiGroup)							//注册用户路由
	router.InitJwtRouter(ApiGroup)							// token操作
	router.InitChatUserRouter(ApiGroup)						// 聊天用户
	router.InitChatSysNoticeRouter(ApiGroup)						// 聊天系统推送
	router.InitMenuRouter(ApiGroup)                  // 注册menu路由
	router.InitAuthorityRouter(ApiGroup)                  // 注册角色路由
	global.GL_LOG.Info("router register success")
	return Router
}