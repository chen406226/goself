package service

import (
	"student/global"
	"student/model/mysqlDb"
	"student/model/request"
)

// @title    CreateSysDictionaryDetail
// @description   create a SysDictionaryDetail
// @param     sysDictionaryDetail               model.SysDictionaryDetail
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSystionaryDetail(sysDictionaryDetail mysqlDb.SysDictionaryDetail) (err error) {
	err = global.GL_DB.Create(&sysDictionaryDetail).Error
	return
}

// @title    DeleteSysDictionaryDetail
// @description   delete a SysDictionaryDetail
// @auth                     （2020/04/05  20:22）
// @param     sysDictionaryDetail               model.SysDictionaryDetail
// @return                    error

func DeleteSysDictionaryDetail(sysDictioanryDetail mysqlDb.SysDictionaryDetail) (err error) {
	err = global.GL_DB.Delete(&sysDictioanryDetail).Error
	return
}

// @title    UpdateSysDictionaryDetail
// @description   update a SysDictionaryDetail
// @param     sysDictionaryDetail          *model.SysDictionaryDetail
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSysDictionaryDetail(sysDictionaryDetail *mysqlDb.SysDictionaryDetail) (err error) {
	err = global.GL_DB.Save(sysDictionaryDetail).Error
	return
}

// @title    GetSysDictionaryDetail
// @description   get the info of a SysDictionaryDetail
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SysDictionaryDetail        SysDictionaryDetail

func GetSysDictionaryDetail(id uint) (err error, sysDictionaryDetail mysqlDb.SysDictionaryDetail) {
	err = global.GL_DB.Where("id = ?",id).First(&sysDictionaryDetail).Error
	return
}

// @title    GetSysDictionaryDetailInfoList
// @description   get SysDictionaryDetail list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (err error, list interface{},total int64) {
	limit := info.PageSize
	offset := limit*(info.Page -1)
	db := global.GL_DB.Model(&mysqlDb.SysDictionaryDetail{})
	var sysDictionaryDetails []mysqlDb.SysDictionaryDetail
	if info.Label != "" {
		db = db.Where("label like ?","%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?",info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?",info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?",info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysDictionaryDetails).Error
	return err ,sysDictionaryDetails,total
}
