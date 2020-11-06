package router

import (
	"github.com/gin-gonic/gin"
	"student/api/menu"
	"student/middleware"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes)  {
	MenuRouter := Router.Group("menu").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		MenuRouter.POST("addBaseMenu", menu.AddBaseMenu)           // 新增菜单
	}
	return MenuRouter
}