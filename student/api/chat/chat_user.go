package chat

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"student/global"
	"student/global/response"
	"student/middleware"
	"student/model/mysqlDb"
	"student/model/request"
	resp "student/model/response"
	"student/service"
	chatService "student/service/chat"
	"student/utils"
	"time"
)

// @Tags Chat
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func Register(c *gin.Context)  {
	var R request.ChatRegisterStruct
	c.ShouldBindJSON(&R)
	UserVerify	:= utils.Rules{
		"Username":	{utils.NotEmpty()},
		"Password":	{utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R,UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}

	user := &mysqlDb.ChatUser{Username:R.Username,NickName:R.NickName,Password: R.Password,HeaderImg:R.HeaderImg}
	err, userReturn := service.ChatRegister(*user)
	if err != nil {
		response.Result(response.ERROR, resp.ChatUserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
	} else {
		response.Result(response.SUCCESS,resp.ChatUserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L request.RegisterAndLoginStruct
	_ = c.ShouldBindJSON(&L)
	UserVerify := utils.Rules{
		//"CaptchaId": {utils.NotEmpty()},
		//"Captcha":   {utils.NotEmpty()},
		"Username":  {utils.NotEmpty()},
		"Password":  {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	//if captcha.VerifyString(L.CaptchaId, L.Captcha) {
		U := &mysqlDb.ChatUser{Username: L.Username, Password: L.Password}
		if err, user := service.ChatLogin(U); err != nil {
			response.FailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
		} else {
			tokenNext(c, *user)
		}
	//} else {
	//	response.FailWithMessage("验证码错误", c)
	//}

}

// @Tags Base
// @Summary 搜索用户
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{user:User{}},"msg":"查询成功"}"
// @Router /chat/login [post]
func SearchUserByUserName(c *gin.Context) {
	var L request.SearchUserStruct
	_ = c.ShouldBindJSON(&L)
	UserVerify := utils.Rules{
		"Username":  {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	//if captcha.VerifyString(L.CaptchaId, L.Captcha) {
	U := &mysqlDb.ChatUser{Username: L.Username}
	if err, user := service.SearchChatUserByUsername(*U); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.Result(response.SUCCESS,resp.ChatUserResponse{User: user}, "查询成功", c)
	}
}

// 獲取朋友列表
func GetUserFriendList(c *gin.Context)  {
	var U request.ChatUserFriendStruct
	c.ShouldBindJSON(&U)
	UserVerify	:= utils.Rules{
		"Username":	{utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(U,UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	user := mysqlDb.ChatUser{Username:U.Username}
	err,userNoticeList := chatService.GetChatUserFriendList(user)
	if err != nil {
		response.Result(response.ERROR, resp.ChatUserFriendResponse{FriendList: userNoticeList}, fmt.Sprintf("%v", err), c)
	} else {
		response.Result(response.SUCCESS,resp.ChatUserFriendResponse{FriendList: userNoticeList}, "获取成功", c)
	}
}

func SetUserAddFriend(c *gin.Context)  {
	var userShip request.ChatUserAddFriend
	c.ShouldBindJSON(&userShip)
	UserVerify	:= utils.Rules{
		"Username":	{utils.NotEmpty()},
		"FriendUsername":	{utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(userShip,UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	user1 := mysqlDb.ChatUser{Username:userShip.Username}
	user2 := mysqlDb.ChatUser{Username:userShip.FriendUsername}
	err := chatService.SetChatUserAddFriend(user1,user2)
	if err!= nil {
		response.Result(response.ERROR, map[int]int{}, fmt.Sprintf("%v", err), c)
	}else{
		response.OkWithMessage("添加好友成功",c)
	}
}





// 登陆后签发jwt
func tokenNext(c *gin.Context,user mysqlDb.ChatUser)  {
	j := &middleware.JWT{[]byte(global.GL_CONFIG.JWT.SigningKey)}
	clams := request.CustomClaims{
		UUID:           uuid.UUID{},
		ID:             user.ID,
		NickName:       user.NickName,
		StandardClaims: jwt.StandardClaims{
			NotBefore:	int64(time.Now().Unix()-1000),			// 签名生效时间
			ExpiresAt:	int64(time.Now().Unix()+60*60*24*7),	// 过期时间
			Issuer:		"admin",								//签名发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
	} else {
		if global.GL_CONFIG.System.UseMultipoint {
			var loginJwt mysqlDb.JwtBlacklist
			loginJwt.Jwt = token
			err, jwtStr := service.GetRedisJWT(user.Username)
			if err == redis.Nil {
				err2 := service.SetRedisJWT(loginJwt, user.Username)
				if err2 != nil {
					response.FailWithMessage("设置登录状态失败",c)
				} else {
					response.OkWithData(resp.ChatLoginResponse{
						User:      user,
						Token:     token,
						ExpiresAt: clams.ExpiresAt * 1000,
					}, c)
				}
			} else if err != nil {
				response.FailWithMessage(fmt.Sprintf("%v", err),c)
			} else {
				var blackJWT mysqlDb.JwtBlacklist
				blackJWT.Jwt = jwtStr
				err3 := service.JsonInBlacklist(blackJWT)
				if err3 != nil {
					response.FailWithMessage("jwt作废失败", c)
				} else {
					err2 := service.SetRedisJWT(loginJwt,user.Username)
					if err2 != nil {
						response.FailWithMessage("设置登录状态失败", c)
					} else {
						response.OkWithData(resp.ChatLoginResponse{
							User:      user,
							Token:     loginJwt.Jwt,
							ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
						}, c)
					}
				}
			}
		} else {
			response.OkWithData(resp.ChatLoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
			}, c)
		}
	}
}