package api

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

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
func CreateApi(c *gin.Context) {
	var A mysqlDb.SysApi
	_ = c.ShouldBindJSON(&A)
	ApiVerify := utils.Rules{
		"Path":					{utils.NotEmpty()},
		"Description":	{utils.NotEmpty()},
		"ApiGroup":			{utils.NotEmpty()},
		"Method":				{utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(A,ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(),c)
		return
	}
	err := service.CreateApi(A)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v",err),c)
	} else {
		response.OkWithMessage("创建成功",c)
	}
}

// @Tags SysApi
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "删除api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
func DeleteApi(c *gin.Context)  {
	var A mysqlDb.SysApi
	_ = c.ShouldBindJSON(&A)
	ApiVerify := utils.Rules{"ID": {utils.NotEmpty()}}
	ApiVerifyErr := utils.Verify(A,ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(),c)
		return
	}
	err := service.DeleteApi(A)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v",err),c)
	} else {
		response.OkWithMessage("删除成功",c)
	}

}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/updateApi [post]
func UpdateApi(c *gin.Context)  {
	var A mysqlDb.SysApi
	_ = c.ShouldBindJSON(&A)
	ApiVerify := utils.Rules{
		"Path":					{utils.NotEmpty()},
		"Description":	{utils.NotEmpty()},
		"ApiGroup":			{utils.NotEmpty()},
		"Method":				{utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(A,ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(),c)
		return
	}
	err := service.UpdateApi(A)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v",err),c)
	} else {
		response.OkWithMessage("更新成功",c)
	}
}

// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func GetApiById(c *gin.Context)  {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo,utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(),c)
		return
	}
	err,api := service.GetApiById(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败,%v",err),c)
	} else {
		//response.OkWithData(api,c)
		response.OkWithData(resp.SysAPIResponse{Api: api},c)
	}
}

// 条件搜索后端看此api

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]

func GetApiList(c *gin.Context)  {
	// 此结构体仅本方法使用
	var sp request.SearchApiParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo,utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(),c)
		return
	}
	err, list, total := service.GetAPIInfoList(sp.SysApi,sp.PageInfo,sp.OrderKey,sp.Desc)
	if err != nil {
		response.FailWithMessage( fmt.Sprintf("获取api列表失败，%v",err),c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     sp.PageInfo.Page,
			PageSize: sp.PageInfo.PageSize,
		},c)
	}
}

// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func GetAllApis(c *gin.Context)  {
	err,apis := service.GetAllApis()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v",err),c)
	} else {
		response.OkWithData(resp.SysAPIListResponse{Apis: apis},c)
	}



}


































