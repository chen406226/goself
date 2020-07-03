package initialize

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"student/global"
)

//初始化数据库并产生数据库全局变量
func Mysql()  {
	mysqlConfig := global.GL_CONFIG.Mysql
	if db, err := gorm.Open("mysql", mysqlConfig.Username+":"+mysqlConfig.Password+"@("+mysqlConfig.Path+")/"+mysqlConfig.Dbname+"?"+mysqlConfig.Config); err != nil {
		global.GL_LOG.Error("Mysql 启动异常", err)
		os.Exit(0) //启动异常停止服务
	} else {
		global.GL_DB = db
		global.GL_DB.DB().SetMaxIdleConns(mysqlConfig.MaxIdleConns)
		global.GL_DB.DB().SetMaxOpenConns(mysqlConfig.MaxOpenConns)
		global.GL_DB.LogMode(mysqlConfig.LogMode)
	}

}