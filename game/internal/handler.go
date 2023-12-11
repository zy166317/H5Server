package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
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

// 闯关逻辑
func initGame() {
	handler(&pb.ChallengeChapterReq{}, handleChallengeChapter)
	handler(&pb.UsePropsReq{}, handleUseProps)
}

func handleSocketConn(args []interface{}) {
	// 收到的 socket连接消息
	req := args[0].(*pb.SocketPingReq)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("conn remote addr", a.RemoteAddr())
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

func handleUseProps(args []interface{}) {
	req := args[0].(*pb.UsePropsReq)
	a := args[1].(gate.Agent)
	p := a.UserData().(*api.Player)
	api.CMgr.UseProps(p, req)
}
