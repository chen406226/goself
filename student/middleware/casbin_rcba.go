package middleware

import (
	"github.com/gin-gonic/gin"
	"student/global"
	"student/global/response"
	"student/model/request"
	"student/service"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		//获取请求的URL
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户角色
		sub := waitUse.AuthorityId
		e := service.Casbin()
		//判断策略中是否存在
		success,_ := e.Enforce(sub, obj, act)
		if global.GL_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			response.Result(response.ERROR, gin.H{}, "权限不足",c)
			c.Abort()
			return
		}
	}
}