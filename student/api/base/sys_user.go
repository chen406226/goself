package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"student/global/response"
	"student/model/mysqlDb"
	"student/model/request"
	resp "student/model/response"
	"student/service"
	"student/utils"
)

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]

func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	UserVerify := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	user := &mysqlDb.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	err, userReturn := service.Register(*user)
	if err != nil {
		response.Result(response.ERROR, resp.SysUserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
	} else {
		response.Result(response.SUCCESS,resp.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context)  {
	var L request.RegisterAndLoginStruct
	_ = c.ShouldBindJSON(&L)
	UserVerify := utils.Rules{
		"Username":		{utils.NotEmpty()},
		"Password":		{utils.NotEmpty()},
		//"CaptchaId":	{utils.NotEmpty()},
		//"Captcha":		{utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	//if captcha.VerifyString(L.CaptchaId, L.Captcha) {
		U := &mysqlDb.SysUser{Username:L.Username,Password:L.Password}
		if err, user := service.Login(U);err != nil {
			response.FailWithMessage(fmt.Sprintf("用户密码错误或%v",err),c)
		} else {
			fmt.Println(user)
			//tokenNext(c, *user)
		}
	//} else {
	//	response.FailWithMessage("验证码错误", c)
	//}


}
