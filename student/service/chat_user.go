package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"student/global"
	"student/model/mysqlDb"
	"student/utils"
)

// @title    Register
// @description   register, 用户注册
// @auth                     （2020/04/05  20:22）
// @param     u               model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func ChatRegister(u mysqlDb.ChatUser) (err error,userInter mysqlDb.ChatUser) {
	var user mysqlDb.ChatUser
	//判断用户是否注册
	notRegister := global.GL_DB.Where("username = ?",u.Username).First(&user).RecordNotFound()
	if !notRegister {
		return errors.New("用户名已注册"), userInter
	} else {
		//附加uuid 密码md5加密
		u.Password = utils.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = global.GL_DB.Create(&u).Error
	}
	return err,u
}


func ChatLogin(u *mysqlDb.ChatUser) (err error, userInter *mysqlDb.ChatUser) {
	var user mysqlDb.ChatUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GL_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return err,&user
	}
	return err, &user
}