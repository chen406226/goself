package service

import (
	"student/global"
	"student/model/mysqlDb"
	"student/model/request"
)

// @title    CreateSysOperationRecord
// @description   create a SysOperationRecord
// @param     sysOperationRecord               model.SysOperationRecord
// @auth                     （2020/04/05  20:22）
// @return    err             error
func CreateSysOperationRecord(syaOperationRecord mysqlDb.SysOperationRecord) error {
	err := global.GL_DB.Create(&syaOperationRecord).Error
	return err
}

// @title    DeleteSysOperationRecord
// @description   delete SysOperationRecords
// @auth                     （2020/04/05  20:22）
// @param     sysOperationRecord               request.IdsReq
// @return                    error

func DeleteSysOperationRecord(sysOperationRecord mysqlDb.SysOperationRecord) error {
	err := global.GL_DB.Delete(&sysOperationRecord).Error
	return err
}

// @title    DeleteSysOperationRecord
// @description   delete SysOperationRecords
// @auth                     （2020/04/05  20:22）
// @param     sysOperationRecord               request.IdsReq
// @return                    error

func DeleteSysOperationRecordByIds(ids request.IdsReq) error {
	//err = global.GVA_DB.Delete(&[]model.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	err := global.GL_DB.Delete(mysqlDb.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    GetSysOperationRecord
// @description   get the info of a SysOperationRecord
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SysOperationRecord        SysOperationRecord

func GetSysOperationRecord(id uint) (err error, sysOperationRecord mysqlDb.SysOperationRecord) {
	err = global.GL_DB.Where("id = ?",id).First(&sysOperationRecord).Error
	return
}

// @title    GetSysOperationRecordInfoList
// @description   get SysOperationRecord list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := limit * (info.Page - 1)
	db := global.GL_DB.Model(&mysqlDb.SysOperationRecord{})
	var sysOperationRecords []mysqlDb.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%" + info.Path + "%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(sysOperationRecords).Error
	return err, sysOperationRecords, total
}
