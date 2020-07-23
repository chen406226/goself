package response

import (
	"github.com/jinzhu/gorm"
	"student/model/mysqlDb"
)

//获得专有用户推送
type GetChatSysNoticeStruct struct {
	NoticeList	[]mysqlDb.ChatSystemNotice	`json:"noticeList"` // 因为mysql存储的是字符串 Annex
	//NoticeList	[]ChatSystemNotice	`json:"noticeList"`
}

//
type ChatSystemNotice struct {
	gorm.Model
	Title	string			`json:"title"`
	Content	string			`json:"content"`
	Annex	interface{}		`json:"annex"`		// 附件存储json 字符串
	Type	int				`json:"type"`		// 8申请添加好友【推送给个人】
	ToAims	string			`json:"toAims"`		// 作用目标 结合type使用ToUsername
}