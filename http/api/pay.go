package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PaySuccessReq 定义支付成功Req
type PaySuccessReq struct {
	Uid string `json:"uid"`
}

// PaySuccessRsp 定义支付成功Rsp
type PaySuccessRsp struct {
	Code int `json:"code"`
	//其余字段根据要求设定
}

func HandlePayReq(c *gin.Context) {
	//方法一 传参方式为Json
	//var req PaySuccessReq
	//err := c.ShouldBindJSON(&req)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, &PaySuccessRsp{Code: 400})
	//	return
	//}
	//方法二  传参方式为form-data
	req1 := &PaySuccessReq{}
	err1 := c.ShouldBind(req1)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, &PaySuccessRsp{Code: 401})
		return
	}
	if req1.Uid == "1" {
		c.JSON(http.StatusOK, &PaySuccessRsp{Code: 200})
	}
}
