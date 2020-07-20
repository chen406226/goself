package mysqlDb

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type ChatUser struct {
	gorm.Model
	UUID 		uuid.UUID		`json:"uuid"`
	Username	string			`json:"userName"`
	Password	string			`json:"-"`
	NickName	string			`json:"nickName" gorm:"default:'呆瓜007'"`
	HeaderImg	string			`json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
}