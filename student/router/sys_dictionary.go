package router

import (
	"github.com/gin-gonic/gin"
	"student/api/dictionary"
	"student/middleware"
)

func InitSysDictionaryRouter(Router *gin.RouterGroup)  {
	SysDictionaryRouter := Router.Group("sysDictionary").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	//{
		SysDictionaryRouter.POST("createSysDictionary",dictionary.CreactSysDictionary)		// 新建
		SysDictionaryRouter.DELETE("deleteSysDictionary",dictionary.DeleteSysDictionary)	// 删除
		SysDictionaryRouter.PUT("updateSysDictionary",dictionary.UpdateSysDictionary)			// 新建
		SysDictionaryRouter.GET("findSysDictionary",dictionary.FindSysDictionary)					// byId
		SysDictionaryRouter.GET("getSysDictionaryList",dictionary.GetSysDictionaryList)		// 新建
	//}

}




