package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leaf_server/game/api"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)

}

func rpcNewAgent(args []interface{}) {
	_ = args[0].(gate.Agent)
	log.Debug("new agent")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	p := a.UserData().(*api.Player)
	p.SaveData()
	log.Debug("close agent")
}
