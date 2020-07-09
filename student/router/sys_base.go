package router

import (
	"github.com/gin-gonic/gin"
	"student/api/base"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("register", base.Register)  			// 注册后台用户
		BaseRouter.POST("login", base.Login)					// 登录
		BaseRouter.POST("captcha",base.Captcha)					// 生成验证码
		BaseRouter.GET("captcha/:captchaId", base.CaptchaImg)	// 验证码图片获取
	}
	return BaseRouter
}