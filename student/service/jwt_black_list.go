package service

import (
	"student/global"
	model "student/model/mysqlDb"
)

// @title    IsBlacklist
// @description   check if the Jwt is in the blacklist or not, 判断JWT是否在黑名单内部
// @auth                     （2020/04/05  20:22）
// @param     jwt             string
// @param     jwtList         model.JwtBlacklist
// @return    err             error

func IsBlacklist(jwt string,jwtList model.JwtBlacklist) bool {
	isNotFound := global.GL_DB.Where("jwt = ?",jwt).First(&jwtList).RecordNotFound()
	return !isNotFound
}