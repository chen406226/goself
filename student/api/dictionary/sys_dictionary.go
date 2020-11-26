package dictionary

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"student/global/response"
	"student/model/mysqlDb"
	"student/service"
)

// @Tags SysDictionary
// @Summary 创建SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "创建SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/createSysDictionary [post]

func CreateSysDictionary(c *gin.Context)  {
	var sysDictionary mysqlDb.SysDictionary
	_ = c.ShouldBindJSON(sysDictionary)
	err := service.CreateSysDictionary(sysDictionary)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v",err),c)
	} else {
		response.OkWithMessage("创建成功",c)
	}
}

// @Tags SysDictionary
// @Summary 删除SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "删除SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionary/deleteSysDictionary [delete]

func DeleteSysDictionary(c *gin.Context)  {
	var sysDictionary mysqlDb.SysDictionary
	_ = c.ShouldBindJSON(sysDictionary)
	err := service.DeleteSysDictionary(sysDictionary)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v"),c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


























