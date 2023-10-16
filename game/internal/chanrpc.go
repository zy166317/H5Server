package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
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
	log.Debug("close agent")
}
