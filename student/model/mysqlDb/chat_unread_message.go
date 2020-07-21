package mysqlDb

import "github.com/jinzhu/gorm"
/*
	未在线存储未读消息
*/
type ChatUnreadMessage struct {
	gorm.Model
	FromUuid	string	`json:"fromUuid"`
	ToUuid		string	`json:"toUuid"`
	Message		string	`json:"message"`
}