package base

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
	"student/utils"
	"time"
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
		//"NickName":    {utils.NotEmpty()},
		//"AuthorityId": {utils.NotEmpty()},
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
	//UserVerify := utils.Rules{
	//	"Username":		{utils.NotEmpty()},
	//	"Password":		{utils.NotEmpty()},
	//	"CaptchaId":	{utils.NotEmpty()},
	//	"Captcha":		{utils.NotEmpty()},
	//}
	//UserVerifyErr := utils.Verify(L, UserVerify)
	//if UserVerifyErr != nil {
	//	response.FailWithMessage(UserVerifyErr.Error(), c)
	//	return
	//}
	//if captcha.VerifyString(L.CaptchaId, L.Captcha) {
		U := &mysqlDb.SysUser{Username:L.Username,Password:L.Password}
		if err, user := service.Login(U);err != nil {
			response.FailWithMessage(fmt.Sprintf("用户密码错误或%v",err),c)
		} else {
			fmt.Println(user)
			tokenNext(c, *user)
		}
	//} else {
	//	response.FailWithMessage("验证码错误", c)
	//}


}

// 登陆后签发jwt
func tokenNext(c *gin.Context,user mysqlDb.SysUser)  {
	j := &middleware.JWT{[]byte(global.GL_CONFIG.JWT.SigningKey)}
	clams := request.CustomClaims{
		UUID:           uuid.UUID{},
		ID:             user.ID,
		NickName:       user.NickName,
		AuthorityId:    user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore:	int64(time.Now().Unix()-1000),			// 签名生效时间
			ExpiresAt:	int64(time.Now().Unix()+60*60*24*7),	// 过期时间
			Issuer:		"qmPlus",								//签名发行者
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
					response.OkWithData(resp.LoginResponse{
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
						response.OkWithData(resp.LoginResponse{
							User:      user,
							Token:     loginJwt.Jwt,
							ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
						}, c)
					}
				}
			}
		} else {
			response.OkWithData(resp.LoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
			}, c)
		}
	}
}

// @Tags jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func JsonInBlacklist(c *gin.Context)  {
	token := c.Request.Header.Get("x-token")
	modelJwt := mysqlDb.JwtBlacklist{Jwt:token}
	err := service.JsonInBlacklist(modelJwt)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("jwt作废失败，%v",err), c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}