package service

import (
	"errors"
	"gorm.io/gorm"
	"student/global"
	"student/model/mysqlDb"
	"student/model/request"
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

// @title    DeleteSysDictionary
// @description   delete a SysDictionary
// @auth                     （2020/04/05  20:22）
// @param     sysDictionary               model.SysDictionary
// @return                    error

func DeleteSysDictionary(d mysqlDb.SysDictionary) (err error) {
	err = global.GL_DB.Delete(d).Delete(&d.SysDictionaryDetails).Error
	return err
}


// @title    UpdateSysDictionary
// @description   update a SysDictionary
// @param     sysDictionary          *model.SysDictionary
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSysDictionary(d *mysqlDb.SysDictionary) (err error) {
	var dict mysqlDb.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":		d.Name,
		"Type": 	d.Type,
		"Status":	d.Status,
		"Desc": 	d.Desc,
	}
	db := global.GL_DB.Where("id = ?",d.ID).First(&dict)
	if dict.Type == d.Type {
		err = db.Updates(sysDictionaryMap).Error
	} else {
		if (!errors.Is(global.GL_DB.First(&mysqlDb.SysDictionary{},"type = ?",d.Type).Error,gorm.ErrRecordNotFound)) {
			return errors.New("存在相同的type,不允许创建")
		}
		err = db.Updates(sysDictionaryMap).Error
	}
	return err
}

// @title    GetSysDictionary
// @description   get the info of a SysDictionary
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SysDictionary        SysDictionary

func GetSysDictoinary(Type string,Id uint) (err error, sysDictionary mysqlDb.SysDictionary) {
	err = global.GL_DB.Where("type = ? or id = ?",Type,Id).Preload("SysDictionaryDetails").
		First(&sysDictionary).Error
	return
}

// @title    GetSysDictionaryInfoList
// @description   get SysDictionary list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSysDictionaryInfoList(info request.SysDictionarySearch) (err error,list interface{},total int64) {
	limit := info.PageSize
	offset := limit*(info.Page - 1)
	db := global.GL_DB.Model(&mysqlDb.SysDictionary{})
	var sysDictoinarys [] mysqlDb.SysDictionary
	if info.Name != "" {
		db = db.Where("`name` like ?","%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` like ?","%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?",info.Status)
	}
	if info.Desc != "" {
		db = db.Where("`desc` like ?","%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysDictoinarys).Error
	return err, sysDictoinarys, total
}



