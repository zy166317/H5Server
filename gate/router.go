package gate

import (
	"leaf_server/game"
	"leaf_server/login"
	"leaf_server/msg"
	"leaf_server/proto/pb"
)

func init() {
	initLogin()
	initGM()
	initGame()
}

func initLogin() {
	//登录
	msg.ProtocolProcessor.SetRouter(&pb.LoginReq{}, login.ChanRPC)
}

func initGM() {
	msg.ProtocolProcessor.SetRouter(&pb.SocketPingReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.EditorChapterReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.GetAllChapterReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.ChapterDetailByIdReq{}, game.ChanRPC)

}

func initGame() {
	msg.ProtocolProcessor.SetRouter(&pb.ChallengeChapterReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.UsePropsReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.SevenDaysRewardReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.StarsRankListReq{}, game.ChanRPC)
}
