package router

import (
	"github.com/gin-gonic/gin"
	"student/api/dictionary"
	"student/middleware"
)

func InitSysDictionaryDetailRouter(Router *gin.RouterGroup)  {
	SysDictionaryDetailRouter := Router.Group("sysDictionaryDetail").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SysDictionaryDetailRouter.POST("createSysDictionaryDetail",dictionary.CreateSysDictionaryDetail)		//新增
		SysDictionaryDetailRouter.DELETE("deleteSysDictionaryDetail",dictionary.DeleteSysDictionaryDetail)		//新增
		SysDictionaryDetailRouter.PUT("updateSysDictionaryDetail",dictionary.UpdateSysDictionaryDetail)		//新增
		SysDictionaryDetailRouter.GET("findSysDictionaryDetail",dictionary.FindSysDictionaryDetail)		//新增
		SysDictionaryDetailRouter.GET("getSysDictionaryDetailList",dictionary.GetSysDictionaryDetailList)		//新增
	}
}