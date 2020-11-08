package v1

import (
	"fmt"
	"server/global/response"
	"server/model/request"
	resp "server/model/response"
	"server/service"
	"server/utils"
	"github.com/gin-gonic/gin"
)

// @Tags casbin
// @Summary 更改角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "更改角色api权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/UpdateCasbin [post]
func UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	AuthorityIdVerifyErr := utils.Verify(cmr, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		response.FailWithMessage(AuthorityIdVerifyErr.Error(), c)
		return
	}
	err := service.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加规则失败，%v", err), c)
	} else {
		response.OkWithMessage("添加规则成功", c)
	}
}

// @Tags casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "获取权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func GetPolicyPathByAuthorityId(c *gin.Context) {
	var cmr request.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	AuthorityIdVerifyErr := utils.Verify(cmr, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		response.FailWithMessage(AuthorityIdVerifyErr.Error(), c)
		return
	}
	paths := service.GetPolicyPathByAuthorityId(cmr.AuthorityId)
	response.OkWithData(resp.PolicyPathResponse{Paths: paths}, c)
}

