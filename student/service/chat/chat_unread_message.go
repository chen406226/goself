package chat

import (
	"fmt"
	"student/global"
	"student/model/mysqlDb"
)

func SaveMessage(fromU string,toU string,msg string)  {
	M := &mysqlDb.ChatUnreadMessage{
		FromUuid: fromU,
		ToUuid:   toU,
		Message:  msg,
	}
	err := global.GL_DB.Create(M).Error
	if err != nil {
		global.GL_LOG.Error(fmt.Sprintf("保存消息错误cause:%v", err))
	}
}

//
func GetMessageByToUuid(toU string) []mysqlDb.ChatUnreadMessage {
	var Ms []mysqlDb.ChatUnreadMessage
	err := global.GL_DB.Where("to_uuid = ?", toU).Find(&Ms).Error
	if err != nil {
		global.GL_LOG.Error(fmt.Sprintf("获取未读消息错误cause:%v", err))
		return []mysqlDb.ChatUnreadMessage{}
	}
	//for k,v := range Ms {
	//	fmt.Println("消息队列",k,v)
	//}
	return Ms
}
// 删除过期已读消息
func DeleteExpiredMessage(toU string)  {
	M := &mysqlDb.ChatUnreadMessage{}
	// 软删除
	//global.GL_DB.Where("to_uuid = ?", toU).Delete(M)
	// 数据删除
	global.GL_DB.Unscoped().Where("to_uuid = ?", toU).Delete(M)
}