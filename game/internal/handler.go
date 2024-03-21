package internal

import (
	"github.com/name5566/leaf/gate"
	"leaf_server/game/api"
	"leaf_server/proto/pb"
	"reflect"
)

func Init() {
	initGmChapter() //GM设置设置关卡地图
	initGame()      //闯关逻辑
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

// gm关卡设计
func initGmChapter() {
	handler(&pb.SocketPingReq{}, handleSocketConn)
	handler(&pb.EditorChapterReq{}, handleEditorChapter)
	handler(&pb.GetAllChapterReq{}, handleGetAllChapter)
	handler(&pb.ChapterDetailByIdReq{}, handleChapterDetail)
}

// 游戏逻辑
func initGame() {
	handler(&pb.ChallengeChapterReq{}, handleChallengeChapter)
	handler(&pb.UsePropsReq{}, handleUseProps)
	handler(&pb.SevenDaysRewardReq{}, handleSevenDayRewards)
	handler(&pb.StarsRankListReq{}, handleStarsRankList)
	handler(&pb.SwitchBafflePlateReq{}, handleSwitchPlate)
	handler(&pb.GMReq{}, handleGM)
	handler(&pb.GameCollectPropsReq{}, handleCollectGameProps)
	handler(&pb.ShopBuyingReq{}, handleShopBuying)
	handler(&pb.YesterdayRankReq{}, handleYesterdayRank)
	handler(&pb.ReceiveRankRewardReq{}, handleYesterdayRankRewards)
	handler(&pb.WatchAdvRewardsReq{}, handleWatchAdv)
}

func handleSocketConn(args []interface{}) {
	// 收到的 socket连接消息
	req := args[0].(*pb.SocketPingReq)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	a.WriteMsg(&pb.SocketPingRsp{
		Code: int32(pb.Error_No),
		Rsp:  req.Req + 1,
	})
}

func handleEditorChapter(args []interface{}) {
	req := args[0].(*pb.EditorChapterReq)
	// 消息的发送者
	a := args[1].(gate.Agent)
	api.EcMgr.EditorChapterInfo(a, req)
}

func handleGetAllChapter(args []interface{}) {
	// 收到的 socket连接消息
	_ = args[0].(*pb.GetAllChapterReq)
	// 消息的发送者
	a := args[1].(gate.Agent)
	api.EcMgr.GetAllChapterInfo(a)
}

func handleChapterDetail(args []interface{}) {
	req := args[0].(*pb.ChapterDetailByIdReq)
	// 消息的发送者
	a := args[1].(gate.Agent)
	api.EcMgr.ChapterDetail(a, req)
}

func handleChallengeChapter(args []interface{}) {
	req := args[0].(*pb.ChallengeChapterReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.CMgr.ChallengeChapter(p, req)
}

func handleCollectGameProps(args []interface{}) {
	req := args[0].(*pb.GameCollectPropsReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.CMgr.CollectProps(p, req)
}

func handleUseProps(args []interface{}) {
	req := args[0].(*pb.UsePropsReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.CMgr.UseProps(p, req)
}

func handleSevenDayRewards(args []interface{}) {
	_ = args[0].(*pb.SevenDaysRewardReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.PlayerMgr.SevenDaysLogin(p)
}

func handleStarsRankList(args []interface{}) {
	req := args[0].(*pb.StarsRankListReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.PlayerMgr.StarsRankList(p, req)
}

func handleSwitchPlate(args []interface{}) {
	req := args[0].(*pb.SwitchBafflePlateReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.PlayerMgr.SwitchPlate(p, req)
}

func handleGM(args []interface{}) {
	req := args[0].(*pb.GMReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.PlayerMgr.GM(p, req)
}

func handleShopBuying(args []interface{}) {
	req := args[0].(*pb.ShopBuyingReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.PlayerMgr.ShopBuying(p, req)
}

func handleYesterdayRank(args []interface{}) {
	_ = args[0].(*pb.YesterdayRankReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.PlayerMgr.YesterdayRank(p)
}

func handleYesterdayRankRewards(args []interface{}) {
	_ = args[0].(*pb.YesterdayRankReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.PlayerMgr.YesterdayRankRewards(p)
}

func handleWatchAdv(args []interface{}) {
	req := args[0].(*pb.WatchAdvRewardsReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.PlayerMgr.WatchAdv(p, req)
}
