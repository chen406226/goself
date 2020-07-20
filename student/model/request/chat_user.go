package request

// 聊天用户注册
type ChatRegisterStruct struct {
	Username	string	`json:"userName"`
	Password	string	`json:"passWord"`
	NickName	string	`json:"nickName" gorm:"default:'呆瓜007'"`
	HeaderImg	string	`json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
}

//聊天用户登录
type ChatLoginStruct struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}