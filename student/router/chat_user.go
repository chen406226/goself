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

	//FilterUserRouter := UserRouter.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	FilterUserRouter := UserRouter.Use(middleware.JWTAuth())
	{
		FilterUserRouter.POST("searchUserByUsername",chat.SearchUserByUserName)	// 搜索用户
		FilterUserRouter.POST("chatUserAddFriend",chat.SetUserAddFriend)	// 添加好友
		FilterUserRouter.POST("getChatUserFriendList",chat.GetUserFriendList)	// 获取朋友列表
		FilterUserRouter.POST("changePassword",)	// 修改密码
		FilterUserRouter.POST("uploadHeaderImg",)	// 上传头像
		FilterUserRouter.DELETE("deleteUser",)		// 删除用户
	}
}