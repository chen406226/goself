package service

import (
	"errors"
	"gorm.io/gorm"
	"student/global"
	"student/model/mysqlDb"
	"student/model/request"
)

// @title    CreateApi
// @description   create base apis, 新增基础api
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @return                    error
func CreateApi(api mysqlDb.SysApi) error {
	if !errors.Is(global.GL_DB.Where("path = ? AND method =?",api.Path,api.Method).First(&mysqlDb.SysApi{}).Error,gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GL_DB.Create(&api).Error
}

// @title    DeleteApi
// @description   delete a base api, 删除基础api
// @param     api             model.SysApi
// @auth                     （2020/04/05  20:22）
// @return                    error
func DeleteApi(api mysqlDb.SysApi) error {
	err := global.GL_DB.Delete(&api).Error
	ClearCasbin(1,api.Path,api.Method)
	return err
}

// @title    UpdateApi
// @description   update a base api, update api
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @return                    error

func UpdateApi(api mysqlDb.SysApi) error {
	var oldA mysqlDb.SysApi
	err := global.GL_DB.Where("id = ?",api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GL_DB.
			Where("path = ? AND method = ?", api.Path, api.Method).
			First(&mysqlDb.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil{
		return err
	} else {
		err = UpdateCasbinApi(oldA.Path,api.Path,oldA.Method,api.Method)
		if err != nil{
			return err
		} else {
			err = global.GL_DB.Save(api).Error
		}
	}
	return err
}

// @title    GetApiById
// @description   根据id获取api
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @param     id              float64
// @return                    error

func GetApiById(id int64) (err error, api mysqlDb.SysApi) {
	err = global.GL_DB.Where("id = ?",id).First(&api).Error
	return
}

// @title    GetAllApis
// @description   get all apis, 获取所有的api
// @auth                     （2020/04/05  20:22）
// @return       err          error
// @return       apis         []SysApi

func GetAllApis() (err error,apis []mysqlDb.SysApi) {
	err = global.GL_DB.Find(&apis).Error
	return
}

// @title    GetAPIInfoList
// @description   get apis by pagination, 分页获取数据
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @param     info            request.PageInfo
// @param     order           string
// @param     desc            bool
// @return    err             error
// @return    list            interface{}
// @return    total           int

func GetAPIInfoList(api mysqlDb.SysApi,info request.PageInfo,order string,desc bool) (err error,list interface{},total int64) {
	limit := info.PageSize
	offset := limit*(info.Page-1)
	db := global.GL_DB.Model(&mysqlDb.SysApi{})
	var apiList []mysqlDb.SysApi
	if api.Path != "" {
		db = db.Where("path LIKE ?","%"+api.Path+"%")
	}
	if api.Description != "" {
		db = db.Where("description LIKE ?","%"+api.Description+"%")
	}
	if api.Method != "" {
		db = db.Where("method = ?",api.Method)
	}
	if api.ApiGroup != "" {
		db = db.Where("api_group = ?",api.ApiGroup)
	}
	err = db.Count(&total).Error
	if err != nil{
	    return err,apiList,total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var orderStr string
			if desc {
				orderStr = order + " desc"
			} else {
				orderStr = order
			}
			err = db.Order(orderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err,apiList,total
}
