package service

import (
	"student/global"
	"student/model/mysqlDb"
)


// @title    SetMenuAuthority
// @description   菜单与角色绑定
// @auth                     （2020/04/05  20:22）
// @param     auth            *model.SysAuthority
// @return                    error
func SetMenuAuthority(auth *mysqlDb.SysAuthority) error {
	var s mysqlDb.SysAuthority
	global.GL_DB.Preload("SysBaseMenus").First(&s,"authority_id = ?",auth.AuthorityId)
	err := global.GL_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus).Error
	return err
}


















