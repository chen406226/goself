package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"student/global"
	"student/model/mysqlDb"
)

// @title    UpdateBaseMenu
// @description   更新路由
// @auth                     （2020/04/05  20:22）
// @param     menu            model.SysBaseMenu
// @return    err             errorgetMenu
func UpdateBaseMenu(menu mysqlDb.SysBaseMenu) (err error) {
	var oldMenu mysqlDb.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	db := global.GL_DB.Where("id = ?",menu.ID).Find(&oldMenu)
	if oldMenu.Name != menu.Name {
		if !errors.Is(global.GL_DB.Where("id <> ? AND name = ?",menu.ID,menu.Name).First(&mysqlDb.SysBaseMenu{}).Error,gorm.ErrRecordNotFound) {
			global.GL_LOG.Debug("存在相同name修改失败")
			return errors.New("存在相同name修改失败")
		}
	}
	err = global.GL_DB.Delete(&mysqlDb.SysBaseMenuParameter{},"sys_base_menu_id = ?",menu.ID).Error
	err = db.Updates(upDateMap).Error
	return err
}

// @title    GetBaseMenuById
// @description   get current menus, 返回当前选中menu
// @auth                     （2020/04/05  20:22）
// @param     id              float64
// @return    err             error
func GetBaseMenuById(id float64) (err error,menu mysqlDb.SysBaseMenu) {
	err = global.GL_DB.Preload("Parameters").Where("id = ?",id).First(&menu).Error
	return
}

// @title    DeleteBaseMenu
// @description   删除基础路由
// @auth                     （2020/04/05  20:22）
// @param     id              float64
// @return    err             error

func DeleteBaseMenu(id float64) (err error) {
	err = global.GL_DB.Preload("Parameters").Where("parent_id = ?",id).First(&mysqlDb.SysBaseMenu{}).Error
	if err != nil{
		var menu mysqlDb.SysBaseMenu
		db := global.GL_DB.Preload("SysAuthoritys").Where("id = ?",id).First(&menu).Delete(&menu)
		err = global.GL_DB.Delete(&mysqlDb.SysBaseMenuParameter{},"sys_base_menu_id = ?",id).Error
		if len(menu.SysAuthoritys) >0 {
			err = global.GL_DB.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys).Error
		} else {
			err = db.Error
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}















