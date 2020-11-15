package router

import (
	"github.com/gin-gonic/gin"
	"student/api/authority"
	"student/middleware"
)

func InitAuthorityRouter(Router *gin.RouterGroup)  {
	AuthorityRouter := Router.Group("authority").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		AuthorityRouter.POST("getAuthorityList",authority.GetAuthorityList)	// 获取角色列表
		AuthorityRouter.POST("createAuthority",authority.CreateAuthority)		// 创建角色
		AuthorityRouter.POST("deleteAuthority",authority.DeleteAuthority)		// 删除角色
		AuthorityRouter.POST("updateAuthority",authority.UpdateAuthority)		// 更新角色
		AuthorityRouter.POST("copyAuthority",authority.CopyAuthority)				// 复制角色
		AuthorityRouter.POST("setDataAuthority",authority.SetDataAuthority)		// 设置角色资源权限
	}
}