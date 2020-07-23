package chat

import (
	"fmt"
	"student/global"
	"student/model/mysqlDb"
)

func SaveChatSystemNotice(n mysqlDb.ChatSystemNotice) error {
	N := &mysqlDb.ChatSystemNotice{
		Title:   n.Title,
		Content: n.Content,
		Annex:   n.Annex,
		Type:    n.Type,
		ToAims:  n.ToAims,
	}
	err := global.GL_DB.Create(N).Error
	if err != nil {
		global.GL_LOG.Error(fmt.Sprintf("SaveChatSystemNoticeErro[cause]:%v", err))
		return err
	}
	return nil
}

func GetChatSystemNoticeByToAims(userNotice mysqlDb.ChatSystemNotice) (err error,noticeList []mysqlDb.ChatSystemNotice) {
	var notices []mysqlDb.ChatSystemNotice
	err = global.GL_DB.Where("to_aims = ?", userNotice.ToAims).Find(&notices).Error
	if err != nil {
		global.GL_LOG.Error(fmt.Sprintf("GetChatSystemNoticeByToAimsError[cause]:%v", err))
		return err, noticeList
	}
	return nil, notices
}

func DeleteChatSystemNoticeByToAims(toAims string) {
	N := &mysqlDb.ChatSystemNotice{}
	// 软删除
	global.GL_DB.Where("to_aims = ?", toAims).Delete(N)
	// 数据删除
	//global.GL_DB.Unscoped().Where("to_aims = ?", toAims).Delete(N)
}