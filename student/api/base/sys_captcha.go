package base

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"student/global"
	"student/global/response"
	resModel "student/model/response"
	"student/utils"
)

func init() {
	go func() {
		//s := captcha.NewMemoryStore(100, time.Duration(global.GL_CONFIG.Captcha.Expiration)*time.Second)
		//fmt.Println("newStorePointer",&s)
		//captcha.SetCustomStore(s)
	}()
}

// @Tags base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context)  {
	captchaId := captcha.NewLen(global.GL_CONFIG.Captcha.KeyLong)
	response.Result(response.SUCCESS,resModel.SysCaptchaResponse{
		CaptchaId:	captchaId,
		PicPath:	"/base/captcha/" + captchaId + ".png",
	},"验证码获取成功",c)
}

// @Tags base
// @Summary 生成验证码图片路径
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha/:captchaId [get]
func CaptchaImg(c *gin.Context)  {
	utils.GinCaptchaServeHTTP(c.Writer, c.Request)
}