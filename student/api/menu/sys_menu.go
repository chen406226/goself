package menu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"student/global/response"
	"student/model/mysqlDb"
	"student/model/request"
	resp "student/model/response"
	"student/service"
	"student/utils"
)

// @Tags menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "新增菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addBaseMenu [post]
//{
//"ID": 0,
//"path": "admin",
//"name": "superAdmin",
//"hidden": false,
//"parentId": "0",
//"component": "views/Admin",
//"meta": {
//"title": "超级管理员",
//"icon": "user-solid",
//"defaultMenu": false,
//"keepAlive": false
//}
//}
func AddBaseMenu(c *gin.Context) {
	var R mysqlDb.SysBaseMenu
	_ = c.ShouldBindJSON(&R)
	MenuVerify := utils.Rules{
		"Path":    {utils.NotEmpty()},
		"ParentId":    {utils.NotEmpty()},
		"Name":    {utils.NotEmpty()},
		"Component":    {utils.NotEmpty()},
		"Sort":    {utils.Ge("0")},
	}
	MenuVerifyErr := utils.Verify(R, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	MetaVerify := utils.Rules{
		"Title":	{utils.NotEmpty()},
	}
	MetaVerifyErr := utils.Verify(R.Meta, MetaVerify)
	if MetaVerifyErr != nil {
		response.FailWithMessage(MetaVerifyErr.Error(), c)
		return
	}
	err := service.AddBaseMenu(R)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败，%v", err),c)
	} else {
		response.OkWithMessage("添加成功", c)
	}

}

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context)  {
	claims,_ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	err, menus := service.GetMenuTree(waitUse.AuthorityId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败, %v",err),c)
	} else {
		response.OkWithData(resp.SysMenusResponse{Menus: menus}, c)
	}

}