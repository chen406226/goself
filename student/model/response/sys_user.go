package response

import "student/model/mysqlDb"

type SysUserResponse struct {
	User	mysqlDb.SysUser	`json:'user'`
}

type LoginResponse struct {
	User		mysqlDb.SysUser	`json:'user'`
	Token		string			`json:'token'`
	ExpiresAt	int64			`json:'expiresAt'`
}