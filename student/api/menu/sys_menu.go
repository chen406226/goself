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

// @Tags menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取基础menu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr !=nil {
		response.FailWithMessage(PageVerifyErr.Error(),c)
	}
	err, menuList, total := service.GetMenuList()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败,%v",err),c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// @Tags menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "更新菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/updateBaseMenu [post]
func UpdateBaseMenu(c *gin.Context)  {
	var menu mysqlDb.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	MenuVerify := utils.Rules{
		"Path":		{"notEmpty"},
		"ParentId":	{utils.NotEmpty()},
		"Name":	{utils.NotEmpty()},
		"Component":	{utils.NotEmpty()},
		"Sort":	{utils.NotEmpty()},
	}
	MenuVerifyErr := utils.Verify(menu,MenuVerify)
	if MenuVerifyErr != nil{
		 response.FailWithMessage(MenuVerifyErr.Error(),c)
		return
	}
	MetaVerify := utils.Rules{
		"Title":	{utils.NotEmpty()},
	}
	MetaVerifyErr := utils.Verify(menu.Meta, MetaVerify)
	if MetaVerifyErr != nil{
	    response.FailWithMessage(MetaVerifyErr.Error(),c)
	    return
	}
	err := service.UpdateBaseMenu(menu)
	if err != nil{
	    response.FailWithMessage(fmt.Sprintf("修改失败：%v",err),c)
	} else {
		response.OkWithMessage("修改成功",c)
	}

}















































