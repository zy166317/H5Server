package msg

import (
	"github.com/name5566/leaf/network/protobuf"
	"leaf_server/proto/pb"
)

var ProtocolProcessor = protobuf.NewProcessor() //protocol processor

// 注册模块消息
func init() {
	initGmChapter()        //Gm设置关卡
	initLogin()            //登录模块
	initChallengeChapter() //闯关模块
}

//定义不同的函数初始不同功能模块的消息

// gm设置关卡
func initGmChapter() {
	ProtocolProcessor.Register(&pb.SocketPingReq{})
	ProtocolProcessor.Register(&pb.SocketPingRsp{})
	ProtocolProcessor.Register(&pb.EditorChapterReq{})
	ProtocolProcessor.Register(&pb.EditorChapterRsp{})
	ProtocolProcessor.Register(&pb.GetAllChapterReq{})
	ProtocolProcessor.Register(&pb.GetAllChapterRsp{})
	ProtocolProcessor.Register(&pb.ChapterDetailByIdReq{})
	ProtocolProcessor.Register(&pb.ChapterDetailByIdRsp{})
}

// 登录模块
func initLogin() {
	ProtocolProcessor.Register(&pb.LoginReq{})
	ProtocolProcessor.Register(&pb.LoginRsp{})
}

// 游戏模块
func initChallengeChapter() {
	ProtocolProcessor.Register(&pb.ChallengeChapterReq{})
	ProtocolProcessor.Register(&pb.ChallengeChapterRsp{})
	ProtocolProcessor.Register(&pb.UsePropsReq{})
	ProtocolProcessor.Register(&pb.UsePropsRsp{})
	ProtocolProcessor.Register(&pb.SevenDaysRewardReq{})
	ProtocolProcessor.Register(&pb.SevenDaysRewardRsp{})
	ProtocolProcessor.Register(&pb.NoticePropsChange{})
	ProtocolProcessor.Register(&pb.StarsRankListReq{})
	ProtocolProcessor.Register(&pb.StarsRankListRsp{})
}
