package msg

import (
	"github.com/name5566/leaf/network/protobuf"
	"leaf_server/proto/pb"
)

var ProtocolProcessor = protobuf.NewProcessor() //protocol processor

func init() {
	ProtocolProcessor.Register(&pb.SocketPingReq{})
	ProtocolProcessor.Register(&pb.SocketPingRsp{})
	ProtocolProcessor.Register(&pb.EditorChapterReq{})
	ProtocolProcessor.Register(&pb.EditorChapterRsp{})
	ProtocolProcessor.Register(&pb.GetAllChapterReq{})
	ProtocolProcessor.Register(&pb.GetAllChapterRsp{})
	ProtocolProcessor.Register(&pb.ChapterDetailByIdReq{})
	ProtocolProcessor.Register(&pb.ChapterDetailByIdRsp{})
	ProtocolProcessor.Register(&pb.LoginReq{})
	ProtocolProcessor.Register(&pb.LoginRsp{})
	ProtocolProcessor.Register(&pb.ChallengeChapterReq{})
	ProtocolProcessor.Register(&pb.ChallengeChapterRsp{})
	ProtocolProcessor.Register(&pb.UsePropsReq{})
	ProtocolProcessor.Register(&pb.UsePropsRsp{})
}

//定义不同的函数初始不同功能模块的消息
