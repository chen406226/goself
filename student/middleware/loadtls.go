package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLHost:	"localhost:443",
			SSLRedirect:true,
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			//如果出现错误，不继续
			fmt.Println(err)
			return
		}
		c.Next()
	}
}