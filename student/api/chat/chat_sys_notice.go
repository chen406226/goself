package chat

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"student/global/response"
	"student/model/mysqlDb"
	"student/model/request"
	resp "student/model/response"
	chatService "student/service/chat"
	"student/utils"
)


func GetUserNotice(c *gin.Context)  {
	var NR request.GetChatSysNoticeStruct
	c.ShouldBindJSON(&NR)
	UserVerify	:= utils.Rules{
		"toAims":	{utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(NR,UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userNotice := mysqlDb.ChatSystemNotice{ToAims:NR.ToAims}
	err,userNoticeList := chatService.GetChatSystemNoticeByToAims(userNotice)
	// dart jsonDecode(v.Annex)
		//for _,v := range userNoticeList {
		//	fmt.Println(v.Annex)
		//	mt := map[string]interface{}{}
		//	json.Unmarshal(utils.StringToByte(v.Annex),v.Annex)
		//	v.Annex = mt
		//	fmt.Println("ksdfjlsdf",mt,mt["userName"])
		//}
	if err != nil {
		response.Result(response.ERROR, resp.GetChatSysNoticeStruct{NoticeList: userNoticeList}, fmt.Sprintf("%v", err), c)
	} else {
		response.Result(response.SUCCESS,resp.GetChatSysNoticeStruct{NoticeList: userNoticeList}, "获取成功", c)
		chatService.DeleteChatSystemNoticeByToAims(NR.ToAims) //请求
	}
}

func SetUserNotice(c *gin.Context)  {
	var NR request.SetChatSysNoticeStruct
	//fmt.Println(c.Request.GetBody())
	c.ShouldBindJSON(&NR)
	fmt.Println(NR)
	UserVerify	:= utils.Rules{
		"title":	{utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(NR,UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	AnnexString,errE :=  json.Marshal(NR.Annex)
	if errE != nil {
		fmt.Println("转字符串报错了")
	}
	userNotice := mysqlDb.ChatSystemNotice{
		Title:   NR.Title,
		Content: NR.Content,
		Annex:   utils.ByteToString(AnnexString),
		Type:    NR.Type,
		ToAims:  NR.ToAims,
	}
	err := chatService.SaveChatSystemNotice(userNotice)
	if err != nil {
		response.Result(response.ERROR, map[string]string{}, fmt.Sprintf("%v", err), c)
	} else {
		response.Result(response.SUCCESS, map[string]string{}, "success", c)
	}
}

//func ()  {
//
//}
//DeleteChatSystemNoticeByToAims
