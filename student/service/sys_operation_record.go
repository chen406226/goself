package service

import (
	"student/global"
	"student/model/mysqlDb"
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