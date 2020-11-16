package service

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"student/global"
	"student/model/mysqlDb"
	"student/model/request"
	"student/utils"
)

// @title    Register
// @description   register, 用户注册
// @auth                     （2020/04/05  20:22）
// @param     u               model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func Register(u mysqlDb.SysUser) (err error, userInter mysqlDb.SysUser) {
	var user mysqlDb.SysUser
	//判断用户是否注册
	notRegister := global.GL_DB.Where("username = ?", u.Username).First(&user).RecordNotFound()
	if !notRegister {
		return errors.New("用户名已注册"), userInter
	} else {
		//附加uuid 密码md5加密
		u.Password = utils.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = global.GL_DB.Create(&u).Error
	}
	return err, u
}

// @title    Login
// @description   login, 用户登录
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @return    err             error
// @return    userInter       *SysUser
func Login(u *mysqlDb.SysUser) (err error, userInter *mysqlDb.SysUser) {
	var user mysqlDb.SysUser
	fmt.Println(u.Password, u.Username)

	u.Password = utils.MD5V([]byte(u.Password))
	fmt.Println(u.Password)
	err = global.GL_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	fmt.Println("err", err)
	if err != nil {
		return err, &user
	}
	err = global.GL_DB.Where("authority_id = ?", user.AuthorityId).First(&user.Authority).Error
	fmt.Println("sdferr", err, user.AuthorityId)
	return err, &user
}

// @title    ChangePassword
// @description   change the password of a certain user, 修改用户密码
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @param     newPassword     string
// @return    err             error
// @return    userInter       *SysUser
func ChangePassword(u *mysqlDb.SysUser, newPsd string) (err error, userInter *mysqlDb.SysUser) {
	var user mysqlDb.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GL_DB.Where("username = ? AND password = ?", u.Username, u.Password).
		First(&user).Update("password", utils.MD5V([]byte(newPsd))).Error
	return err, u
}

// @title    GetUserInfoList
// @description   get user list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int
func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := (info.Page - 1) * limit
	db := global.GL_DB.Model(&mysqlDb.SysUser{})
	var userList []mysqlDb.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error
func SetUserAuthority(uuid uuid.UUID, authorityId string) error {
	err := global.GL_DB.Where("uuid = ?", uuid).First(&mysqlDb.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
func DeleteUser(id float64) (err error) {
	err = global.GL_DB.Where("id = ?", id).Delete(&mysqlDb.SysUser{}).Error
	return
}

// @title    SetUserInfo
// @description   set the authority of a certain user, 设置用户信息
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error
func SetUserInfo(u mysqlDb.SysUser) (err error, user mysqlDb.SysUser) {
	err = global.GL_DB.Updates(&u).Error
	return err, u
}
