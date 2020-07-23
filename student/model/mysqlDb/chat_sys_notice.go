package mysqlDb

import "github.com/jinzhu/gorm"

// 系统推送消息
type ChatSystemNotice struct {
	gorm.Model
	Title	string		`json:"title"`
	Content	string		`json:"content"`
	Annex	string		`json:"annex"`		// 附件存储json 字符串
	Type	int			`json:"type"`		// 8申请添加好友【推送给个人】
	ToAims	string		`json:"toAims"`		// 作用目标 结合type使用ToUsername
}