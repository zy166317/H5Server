package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/name5566/leaf/log"
	"io/ioutil"
	"net/http"
)

// GetUidReq 定义支付成功Req
type GetUidReq struct {
	JsCode string `json:"js_code"` //登陆时code
}

// GetUidRsp 定义支付成功Rsp
type GetUidRsp struct {
	Openid     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	Unionid    string `json:"unionid"`     //开放平台唯一标识
	Errcode    int    `json:"errcode"`     //错误码
	Errmsg     string `json:"errmsg"`      //错误信息
}

func HandleGetUid(c *gin.Context) {
	//方法一 传参方式为Json
	//var req PaySuccessReq
	//err := c.ShouldBindJSON(&req)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, &PaySuccessRsp{Code: 400})
	//	return
	//}
	//方法二  传参方式为form-data
	req1 := &GetUidReq{}
	err := c.ShouldBind(req1)
	if err != nil {
		c.JSON(http.StatusBadRequest, &GetUidRsp{Errcode: 401})
		return
	}
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=wx2b4a75cfbd8a82e3&secret=146bb3cc2c04efa5b3026a8e6535dffe&js_code=" + req1.JsCode + "&grant_type=authorization_code")
	log.Release("error Get", err)
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	log.Release("error Read Body", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, &GetUidRsp{Errcode: 401})
		return
	}
	rsp := &GetUidRsp{}
	json.Unmarshal(all, rsp)
	c.JSON(http.StatusOK, rsp)
}
