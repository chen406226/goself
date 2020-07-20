package response

import "student/model/mysqlDb"

type ChatLoginResponse struct {
	User		mysqlDb.ChatUser	`json:"user"`
	Token		string				`json:"token"`
	ExpiresAt	int64				`json:"expiresAt"`
}

type ChatUserResponse struct {
	User		mysqlDb.ChatUser	`json:"user"`
}