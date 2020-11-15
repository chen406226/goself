package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
	"student/global"
	"student/model/mysqlDb"
	"student/model/request"
	resp "student/model/response"
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

// @title    GetInfoList
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22）
// @param     info            request.PaveInfo
// @return                    error
// 分页获取数据
func GetAuthorityInfoList(info request.PageInfo) (err error,list interface{},total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page -1)
	db := global.GL_DB
	var authority []mysqlDb.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").
		Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		// v不能直接赋值 是拷贝的 指针变了
		for k,_ := range authority {
			err = findChildrenAuthority(&authority[k])
		}
	}
	return err,authority,total
}

// @title    findChildrenAuthority
// @description   查询子角色
// @auth                     （2020/04/05  20:22）
// @param     auth            *model.SysAuthority
// @return                    error
func findChildrenAuthority(authority *mysqlDb.SysAuthority) (err error) {
	err = global.GL_DB.Preload("DataAuthorityId").
		Where("parent_id = ?",authority.AuthorityId).
		Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

// @title    CreateAuthority
// @description   创建一个角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// @return    authority       model.SysAuthority

func CreateAuthority(auth mysqlDb.SysAuthority) (err error,authority mysqlDb.SysAuthority) {
	var authorityBox mysqlDb.SysAuthority
	if !errors.Is(global.GL_DB.Where("authority_id = ?",auth.AuthorityId).First(&authorityBox).Error,gorm.ErrRecordNotFound){
	  return errors.New("存在相同角色Id"),auth
	}
	err = global.GL_DB.Create(&auth).Error
	return err,auth
}

// @title    CopyAuthority
// @description   复制一个角色
// @auth                     （2020/04/05  20:22）
// @param     copyInfo        response.SysAuthorityCopyResponse
// @return                    error
// @return    authority       model.SysAuthority

func CopyAuthority(copyInfo resp.SysAuthorityCopyResponse) (err error,authority mysqlDb.SysAuthority) {
	var authorityBox mysqlDb.SysAuthority
	if !errors.Is(global.GL_DB.Where("authority_id = ?",copyInfo.Authority.AuthorityId).First(&authorityBox).Error,gorm.ErrRecordNotFound){
		return errors.New("存在相同角色id"),authority
	}
	copyInfo.Authority.Children = []mysqlDb.SysAuthority{}
	err,menus := GetMenuAuthority(copyInfo.OldAuthorityId)
	var baseMenu [] mysqlDb.SysBaseMenu
	for _, v := range menus {
		intNum,_ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu,v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.GL_DB.Create(&copyInfo.Authority).Error
	paths := GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = UpdateCasbin(copyInfo.Authority.AuthorityId,paths)
	if err != nil{
		_ = DeleteAuthority(&copyInfo.Authority)
	}
	return err,copyInfo.Authority
}

// @title    DeleteAuthority
// @description   删除角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// 删除角色
func DeleteAuthority(auth *mysqlDb.SysAuthority) (err error) {
	if !errors.Is(global.GL_DB.Where("authority_id = ?",auth.AuthorityId).First(&mysqlDb.SysUser{}).Error,gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GL_DB.Where("parent_id = ?",auth.AuthorityId).First(&mysqlDb.SysAuthority{}).Error,gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色禁止删除")
	}
	db := global.GL_DB.Preload("SysBaseMenus").Where("authority_id = ?",auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if len(auth.SysBaseMenus) > 0 {
		err = global.GL_DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus).Error
		//err = db.Association("SysBaseMenus").Delete(&auth)
	} else {
		err = db.Error
	}
	ClearCasbin(0,auth.AuthorityId)
	return err
}

// @title    SetDataAuthority
// @description   设置角色资源权限
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
func SetDataAuthority(auth mysqlDb.SysAuthority) error {
	var s mysqlDb.SysAuthority
	global.GL_DB.Preload("DataAuthorityId").First(&s,"authority_id = ?",auth.AuthorityId)
	err := global.GL_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId).Error
	return err
}

// @title    UpdateAuthority
// @description   更改一个角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// @return    authority       model.SysAuthority

func UpdateAuthority(auth mysqlDb.SysAuthority) (err error,authority mysqlDb.SysAuthority) {
	err = global.GL_DB.Where("authority_id = ?",auth.AuthorityId).First(&mysqlDb.SysAuthority{}).Updates(&auth).Error
	return err,auth
}









