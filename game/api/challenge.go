package api

import "leaf_server/proto/pb"

type Challenge struct {
}

func (c *Challenge) ChallengeChapter(p *Player, req *pb.ChallengeChapterReq) {
	if int(req.ChapterId) > len(p.CheckPointRecords)+1 {
		p.Agent.WriteMsg(&pb.ChallengeChapterRsp{
			Code: int32(pb.Error_ChapterLimit),
		})
		return
	}
	//根据星级发放奖励 当前没配好表默认加100
	p.DbPlayer.Gold += 100
	//添加或更新关卡星级
	p.CheckPointRecords[req.ChapterId] = req.Stars
	p.Agent.WriteMsg(&pb.ChallengeChapterRsp{
		Code:           int32(pb.Error_No),
		GoldNum:        100,
		CurrGoldNum:    p.DbPlayer.Gold,
		ChapterRecords: p.CheckPointRecords,
	})
}

func (c *Challenge) UseProps(p *Player, req *pb.UsePropsReq) {
	//当前没配置表 忽略判断
	if p.DbPlayer.Gold < 100 {
		p.Agent.WriteMsg(&pb.UsePropsRsp{
			Code: int32(pb.Error_GoldNotEnough),
		})
		return
	}
	p.DbPlayer.Gold -= 100
	p.Agent.WriteMsg(&pb.UsePropsRsp{
		Code:        int32(pb.Error_No),
		CurrGoldNum: p.DbPlayer.Gold,
	})
}
