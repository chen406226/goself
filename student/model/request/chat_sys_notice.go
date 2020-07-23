package request

//获得专有用户推送
type GetChatSysNoticeStruct struct {
	ToAims string	`json:"toAims"`
}

//设置专有用户推送
type SetChatSysNoticeStruct struct {
	Title	string			`json:"title"`
	Content	string			`json:"content"`
	Annex	interface{}		`json:"annex"`		// 附件存储json 字符串
	Type	int				`json:"type"`		// 8申请添加好友【推送给个人】
	ToAims	string			`json:"toAims"`		// 作用目标 结合type使用ToUsername
}