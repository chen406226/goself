package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR	= 7
	SUCCESS = 0
)


type Response struct {
	Code int			`json:"code"`
	Data interface{}	`json:"data"`
	Msg  string			`json:"msg"`
}

func Result(code int,data interface{},msg string,c *gin.Context)  {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context){
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(msg string,c *gin.Context)  {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithData(data interface{},c *gin.Context)  {
	Result(SUCCESS,data,"操作成功",c)
}

func Fail(c *gin.Context){
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(msg string,c *gin.Context)  {
	Result(ERROR, map[string]int{}, msg, c)
}
