package api

import (
	"leaf_server/common"
	"leaf_server/db"
	"leaf_server/gamedata"
	"leaf_server/proto/pb"
)

type Challenge struct {
}

// ChallengeChapter 挑战关卡
func (c *Challenge) ChallengeChapter(p *Player, req *pb.ChallengeChapterReq) {
	if int(req.ChapterId) > len(p.CheckPointRecords)+1 {
		p.Agent.WriteMsg(&pb.ChallengeChapterRsp{
			Code: int32(pb.Error_ChapterLimit),
		})
		return
	}
	index := common.GetCf(common.CHAPTER_STARS).Index(int(req.ChapterId))
	if index == nil {
		p.Agent.WriteMsg(&pb.ChallengeChapterRsp{
			Code: int32(pb.Error_ReadConfigTable),
		})
		return
	}
	chapterStars := index.(*gamedata.Chapter_stars)
	rewards := make(map[int32]int32)
	for k, v := range chapterStars.Star {
		if req.CompleteTime >= int32(v[0]) && req.CompleteTime <= int32(v[1]) {
			rewards[100000] += int32(chapterStars.Pos[k])
			p.CheckPointRecords[req.ChapterId] = int32(k + 1)
			p.DbPlayer.Gold += int32(chapterStars.Pos[k])
			break
		}
	}
	//添加或更新关卡星级
	p.Agent.WriteMsg(&pb.ChallengeChapterRsp{
		Code:           int32(pb.Error_No),
		Rewards:        rewards,
		ChapterRecords: p.CheckPointRecords,
	})
	p.Agent.WriteMsg(&pb.NoticePropsChange{
		Gold:     p.DbPlayer.Gold,
		Backpack: p.BackPack,
	})
	db.UpdatePlayerStars(c.CalcAllStars(p.CheckPointRecords), p.DbPlayer.Uid)
}

// UseProps 使用道具
func (c *Challenge) UseProps(p *Player, req *pb.UsePropsReq) {
	//当前没配置表 忽略判断
	index := common.GetCf(common.ITEM).Index(int(req.PropId))
	if index == nil {
		p.Agent.WriteMsg(&pb.UsePropsRsp{
			Code: int32(pb.Error_ReadConfigTable),
		})
		return
	}
	item := index.(*gamedata.Item)
	if p.DbPlayer.Gold < int32(item.Price[1]) {
		p.Agent.WriteMsg(&pb.UsePropsRsp{
			Code: int32(pb.Error_GoldNotEnough),
		})
		return
	}
	p.DbPlayer.Gold -= int32(item.Price[1])
	p.Agent.WriteMsg(&pb.UsePropsRsp{
		Code:        int32(pb.Error_No),
		CurrGoldNum: p.DbPlayer.Gold,
	})
}

// CalcAllStars 计算当前星星总数
func (c *Challenge) CalcAllStars(records map[int32]int32) int {
	stars := 0
	for _, v := range records {
		stars += int(v)
	}
	return stars
}
