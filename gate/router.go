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
	msg.ProtocolProcessor.SetRouter(&pb.SwitchBafflePlateReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.GMReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.GameCollectPropsReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.ShopBuyingReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.YesterdayRankReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.ReceiveRankRewardReq{}, game.ChanRPC)
	msg.ProtocolProcessor.SetRouter(&pb.WatchAdvRewardsReq{}, game.ChanRPC)

}
