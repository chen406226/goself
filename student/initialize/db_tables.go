package initialize

import (
	"student/global"
	//"student/model"
	model "student/model/mysqlDb"
)

//注册数据库
func DBTables()  {
	db := global.GL_DB
	db.AutoMigrate(
		model.ExaCustomer{},
		model.SysApi{},
		model.SysAuthority{},
		model.SysBaseMenu{},
		model.SysUser{},
		model.SysWorkflow{},
		model.SysWorkflowStepInfo{},
		model.JwtBlacklist{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},
		model.ChatUser{},
	)
	global.GL_LOG.Debug("register db tables success")
}