package main

import (
	//"fmt"
	//uuid "github.com/satori/go.uuid"
	"student/global"
	start "student/init"
	"student/initialize"
)

func main() {
	//fmt.Println("uuid",uuid.NewV1())
	//fmt.Println("uuid4",uuid.NewV4())
	//pdf1.Main()
	//pdf1.Ost()
	switch global.GL_CONFIG.System.DbType {
		case "mysql":
			initialize.Mysql()
		default:
			initialize.Mysql()
	}
	initialize.DBTables()
	go start.RunChatServer()
	defer global.GL_DB.Close()


	// end
	start.RunServer()

	//println(global.Test())
}
