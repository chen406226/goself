package main

import (
	"student/global"
	start "student/init"
	"student/initialize"
)

func main() {
	//fmt.Println("sdf")
	//pdf1.Main()
	//pdf1.Ost()
	switch global.GL_CONFIG.System.DbType {
		case "mysql":
			initialize.Mysql()
		default:
			initialize.Mysql()
	}
	initialize.DBTables()
	defer global.GL_DB.Close()
	start.RunServer()
	//println(global.Test())
}
