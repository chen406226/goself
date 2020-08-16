package chat

import (
	"errors"
	"fmt"
	"student/global"
	"student/model/mysqlDb"
)

func GetChatUserFriendList(user mysqlDb.ChatUser) (err error,userList []mysqlDb.ChatUser) {
	var friends []mysqlDb.ChatUserFriendList
	err = global.GL_DB.Where(&mysqlDb.ChatUserFriendList{Username:user.Username}).Find(&friends).Error
	var users []mysqlDb.ChatUser
	if err!=nil {
		return err,users
	}
	for _,friend := range friends {
		fmt.Println(friend)
		var user mysqlDb.ChatUser
		err = global.GL_DB.Where("username = ?", friend.FriendUsername).Find(&user).Error
		users = append(users, user)
	}

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
	}else{
		v1 := mysqlDb.ChatUserFriendList{Username:u1.Username,FriendUsername:u2.Username}
		v2 := mysqlDb.ChatUserFriendList{Username:u2.Username,FriendUsername:u1.Username}
		err1 := global.GL_DB.Create(&v1).Error
		err2 := global.GL_DB.Create(&v2).Error
		if err1 != nil || err2 !=nil {
			return errors.New("添加失敗")
		}
	}
	return nil
}
