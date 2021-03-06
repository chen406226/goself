package operation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"student/global/response"
	"student/model/mysqlDb"
	"student/model/request"
	resp "student/model/response"
	"student/service"
)

// @Tags SysOperationRecord
// @Summary 创建SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "创建SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/createSysOperationRecord [post]
func CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord mysqlDb.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	err := service.CreateSysOperationRecord(sysOperationRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err),c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "删除SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]

func DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord mysqlDb.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	err := service.DeleteSysOperationRecord(sysOperationRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 批量删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecordByIds [delete]

func DeleteSysOperationRecordByIds(c *gin.Context)  {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	err := service.DeleteSysOperationRecordByIds(ids)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 用id查询SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "用id查询SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOperationRecord/findSysOperationRecord [get]

func FindSysOperationRecord(c *gin.Context) {
	var sysOperationRecord	mysqlDb.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	err, resysOperationRecord := service.GetSysOperationRecord(sysOperationRecord.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resysOperationRecord":resysOperationRecord},c)
	}
}

// @Tags SysOperationRecord
// @Summary 分页获取SysOperationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysOperationRecordSearch true "分页获取SysOperationRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/getSysOperationRecordList [get]

func GetSysOperationRecordList(c *gin.Context) {
	var pageInfo request.SysOperationRecordSearch
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := service.GetSysOperationRecordInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败,%v",err),c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		},c)
	}
}
