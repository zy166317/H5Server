package msg

import (
	"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/network/protobuf"
	"leaf_server/proto/pb"
)

// Processor 使用json注册消息无需关注先后顺序，如使用protocol则需按照先后顺序进行
var Processor = json.NewProcessor()             //json processor
var ProtocolProcessor = protobuf.NewProcessor() //protocol processor

func init() {
	Processor.Register(&pb.TestReq{})
	Processor.Register(&pb.TestRsp{})
}

//定义不同的函数初始不同功能模块的消息

func msgInitChat() {

}

func msgInitBattle() {

}

func msgInitCave() {

}
