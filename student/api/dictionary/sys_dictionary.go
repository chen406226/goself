package dictionary

import "github.com/gin-gonic/gin"

// @Tags SysDictionary
// @Summary 创建SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "创建SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/createSysDictionary [post]

func CreateSysDictionary(c *gin.Context)  {
	var sysDictionary mysqlDb是飞洒
}