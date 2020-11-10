package router

import (
	"github.com/gin-gonic/gin"
	"student/api/menu"
	"student/middleware"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes)  {
	MenuRouter := Router.Group("menu").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		MenuRouter.POST("getMenu",			menu.GetMenu)					// 获取用户动态路由菜单树
		MenuRouter.POST("getMenuList",		menu.GetMenuList)           	// 分页获取基础menu列表
		MenuRouter.POST("addBaseMenu",		menu.AddBaseMenu)           	// 新增菜单
		MenuRouter.POST("updateBaseMenu",	menu.UpdateBaseMenu)           	// 更新菜单
		MenuRouter.POST("getBaseMenuTree",	menu.GetBaseMenuTree)           // 获取用户动态路由
		MenuRouter.POST("deleteBaseMenu",	menu.DeleteBaseMenu)           	// 删除菜单
		MenuRouter.POST("getBaseMenuById",	menu.GetBaseMenuById)           // 获取菜单byId
		MenuRouter.POST("addMenuAuthority",	menu.AddMenuAuthority)          // 增加menu和角色关联关系
		MenuRouter.POST("getMenuAuthority",	menu.GetMenuAuthority)          // 获取指定角色menu
	}
	return MenuRouter
}