package init

import (
	"fmt"
	"net/http"
	"student/global"
	"student/initialize"
	"time"
)

func RunServer()  {
	if global.GL_CONFIG.System.UseMultipoint {
		initialize.Redis()
	}
	Router := initialize.Router()
	Router.Static("/form-generator","./resource/static")  //路由前缀
	println(global.GL_CONFIG.System.Addr)
	address := fmt.Sprintf(":%d", global.GL_CONFIG.System.Addr)
	s := &http.Server{
		Addr:			address,
		Handler:		Router,
		ReadTimeout:	time.Second*10,
		WriteTimeout:	time.Second*10,
		MaxHeaderBytes: 1<<20,
	}

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GL_LOG.Debug("server run success on ", address)
	fmt.Printf(`服务已启动
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, s.Addr)
	global.GL_LOG.Error(s.ListenAndServe())
}
