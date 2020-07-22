package mysqlDb

import "github.com/jinzhu/gorm"

// 系统推送消息
type ChatSystemNotice struct {
	gorm.Model
	Title	string		`json:"title"`
	Content	string		`json:"content"`
	Annex	string		`json:"Annex"` // 附件存储json 字符串
	Type	int			`json:"type"`

}