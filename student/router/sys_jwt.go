package router

import (
	"github.com/gin-gonic/gin"
	"student/api/base"
	"student/middleware"

	//"student/middleware"
)

func InitJwtRouter(Router *gin.RouterGroup)  {
	ApiRouter := Router.Group("jwt").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//ApiRouter := Router.Group("jwt")
	{
		ApiRouter.POST("jsonInBlacklist",base.JsonInBlacklist) //jwt加入黑名单
	}
}
