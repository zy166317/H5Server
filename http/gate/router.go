package gate

import (
	"github.com/gin-gonic/gin"
	"leaf_server/http/api"
)

// PaySuccessInfo eg:接收SDK支付成功通知
func PaySuccessInfo(c *gin.Context) {
	api.HandlePayReq(c)
}
