package initialize

import (
	"github.com/gin-gonic/gin"
	"student/global"
)

func Router() *gin.Engine {
	Router := gin.Default()
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GL_LOG.Debug("use middleware logger")
	// 跨域
	//Router.Use(middleware.Cors())
	//global.GVA_LOG.Debug("use middleware cors")
	ApiGroup := Router.Group("")
	_ = ApiGroup
	global.GL_LOG.Info("router register success")
	return Router
}