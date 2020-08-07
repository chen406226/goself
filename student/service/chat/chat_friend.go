package chat

import (
	"errors"
	"fmt"
	"student/global"
	"student/model/mysqlDb"
)

func GetChatUserFriendList(user mysqlDb.ChatUser) (err error,userList []mysqlDb.ChatUser) {
	var users []mysqlDb.ChatUser
	err = global.GL_DB.Where("username = ?", user.Username).Find(&users).Error
	if err != nil {
		global.GL_LOG.Error(fmt.Sprintf("GetChatUserFriendListError[cause]:%v", err))
		return err, users
	}
	return nil, users
}

func SetChatUserAddFriend(u1 mysqlDb.ChatUser,u2 mysqlDb.ChatUser) (err error) {
	var friendShip mysqlDb.ChatUserFriendList
	notFriendship := global.GL_DB.Where("username = ? AND friend_username = ?", u1.Username, u2.Username).First(&friendShip).RecordNotFound()
	if !notFriendship {
		return errors.New("已是好友")
	}
	return nil
}
