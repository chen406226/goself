package router
/*
	聊天用户相关路由
*/
import (
	"github.com/gin-gonic/gin"
	"student/api/chat"
	"student/middleware"
)

func InitChatUserRouter(Router *gin.RouterGroup)  {
	UserRouter := Router.Group("chat")
	{
		UserRouter.POST("register", chat.Register)		// 用户注册
		UserRouter.POST("login", chat.Login)			// 用户登录
	}

	FilterUserRouter := UserRouter.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		FilterUserRouter.POST("changePassword",)	// 修改密码
		FilterUserRouter.POST("uploadHeaderImg",)	// 上传头像
		FilterUserRouter.DELETE("deleteUser",)		// 删除用户
	}
}