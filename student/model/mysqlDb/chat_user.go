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
	HeaderImg	string			`json:"headerImg" gorm:"default:'https://s1.ax1x.com/2020/07/22/UHigmt.jpg'"`
}

// 朋友列表  可以另一种写法每个添加加入两个  每个人都做一次主宾 好处 修改备注啥的 针对定制信息 坏处 数据量大了两倍
type ChatUserFriendList struct {
	gorm.Model
	Username			string		`json:"username"`
	FriendUsername		string		`json:"friendUsername"`
	FriendRemark		string		`json:"friendRemark"`
	FriendGroup			string		`json:"friendGroup"`
	//FromUUID	string		`json:"fromUuid"`
	//ToUUID		string		`json:"toUuid"`
}