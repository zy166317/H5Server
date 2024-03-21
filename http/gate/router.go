package gate

import (
	"github.com/gin-gonic/gin"
	"leaf_server/http/api"
)

// GetUid eg:获取微信uid
func GetUid(c *gin.Context) {
	api.HandleGetUid(c)
}
