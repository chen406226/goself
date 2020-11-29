package service

import (
	"errors"
	"gorm.io/gorm"
	"student/global"
	"student/model/mysqlDb"
)

// @title    CreateSysDictionary
// @description   create a SysDictionary
// @param     sysDictionary               model.SysDictionary
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSysDictionary(d mysqlDb.SysDictionary) (err error) {
	if (!errors.Is(global.GL_DB.First(&mysqlDb.SysDictionary{},"type = ?",d.Type).Error,gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type,不允许创建")
	} else {
		return global.GL_DB.Create(&d).Error
	}
}











