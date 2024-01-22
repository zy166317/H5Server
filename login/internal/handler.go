package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leaf_server/db"
	"leaf_server/game/api"
	"leaf_server/proto/pb"
	"reflect"
)

func Init() {
	handler(&pb.LoginReq{}, handleLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	req := args[0].(*pb.LoginReq)
	a := args[1].(gate.Agent)
	log.Release("玩家" + req.Uid + "登录")
	info, _ := db.GetPlayerInfo(req.Uid)
	if info != nil {
		api.PlayerMgr.UserLogin(info, a, req)
	} else {
		//新用户走创角
		if req.InviteUid != "" {

		}
		api.PlayerMgr.CreateRole(req, a)
	}
}
