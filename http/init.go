package http

import (
	"github.com/gin-gonic/gin"
	"leaf_server/http/gate"
)

func InitHttpService(addr string) {
	engine := gin.Default()
	//路由分组
	v1 := engine.Group("/pay")
	{
		v1.POST("/Ping", func(context *gin.Context) {
			gate.PaySuccessInfo(context)
		})
	}
	//启动
	engine.Run(addr)
}
