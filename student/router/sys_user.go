package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"student/api/base"
	"student/middleware"
)

func InitUserRouter(Router *gin.RouterGroup)  {
	UserRouter := Router.Group("user").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	fmt.Sprintf("dfs")
	{
		UserRouter.POST("register", base.Register)  							// 注册后台用户
		UserRouter.POST("changePassword",base.ChangePassword)			// 修改密码
		UserRouter.POST("getUserList",base.GetUserList) 					// 分页获取用户列表
		UserRouter.POST("setUserAuthority",base.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser",base.DeleteUser)						// 删除用户
		UserRouter.PUT("setUserInfo",base.SetUserInfo)						// 设置用户信息
	}
}