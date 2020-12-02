package router

import (
	"github.com/gin-gonic/gin"
	"student/api/operation"
	"student/middleware"
)

func InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	SysOperationRecordRouter := Router.Group("sysOperationRecord").
		Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SysOperationRecordRouter.POST("createSysOperationRecord", operation.CreateSysOperationRecord)	// new operationRecord
		SysOperationRecordRouter.DELETE("deleteSysOperationRecord", operation.DeleteSysOperationRecord)	// Delete operationRecord
		SysOperationRecordRouter.DELETE("deleteSysOperationRecordByIds", operation.DeleteSysOperationRecordByIds)	// 批量 Delete operationRecord
		SysOperationRecordRouter.GET("findSysOperationRecord", operation.FindSysOperationRecord)	// get operation by id
		SysOperationRecordRouter.GET("getSysOperationRecordList", operation.GetSysOperationRecordList)	// get operation list
	}
}