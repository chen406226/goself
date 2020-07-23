package router

import (
	"github.com/gin-gonic/gin"
	"student/api/chat"
	"student/middleware"
)

func InitChatSysNoticeRouter(Router *gin.RouterGroup)  {
	SysNoticeRouter := Router.Group("chatSys/notice")
	{
		SysNoticeRouter.POST("getUserNotice", chat.GetUserNotice)		// 获得专有用户推送（添加好友用）
		SysNoticeRouter.POST("setUserNotice", chat.SetUserNotice)			// 设置转悠用户推送（添加好友用）
	}

	FilterUserRouter := SysNoticeRouter.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		FilterUserRouter.POST("changePassword",)	// 修改密码
		FilterUserRouter.DELETE("deleteUser",)		// 删除用户
	}
}