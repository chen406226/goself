package router

import (
	"github.com/gin-gonic/gin"
	"student/api/api"
	"student/middleware"
)

// api router

func InitApiRouter(Router *gin.RouterGroup)  {
	ApiRouter := Router.Group("api").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("createApi",api.CreateApi)		// 创建Api
		ApiRouter.POST("deleteApi",api.DeleteApi)		// 删除Api
		ApiRouter.POST("updateApi",api.UpdateApi)		// 更新Api
		ApiRouter.POST("getApiList",api.GetApiList)	// 获取Api列表
		ApiRouter.POST("getApiById",api.GetApiById)	// 获取单条Api
		ApiRouter.POST("getAllApis",api.GetAllApis)	// 获取所有Api
	}
}