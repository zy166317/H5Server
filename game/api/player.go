package api

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leaf_server/common"
	"leaf_server/db"
	"leaf_server/gamedata"
	"leaf_server/proto/pb"
	"leaf_server/util"
	"strconv"
	"time"
)

type Player struct {
	Agent             gate.Agent
	DbPlayer          *db.Player
	ChallengeInfo     Challenge
	BackPack          map[int32]int32
	CheckPointRecords map[int32]int32
	LoginDay          []int32
}

func NewPlayer(player *db.Player, a gate.Agent) *Player {
	return &Player{
		Agent:    a,
		DbPlayer: player,
	}
}

// LoadData 加载玩家数据库数据
func (p *Player) LoadData() {
	util.ConvertString2Structure(p.DbPlayer.BackPack, &p.BackPack)
	util.ConvertString2Structure(p.DbPlayer.CheckPointRecord, &p.CheckPointRecords)
	util.ConvertString2Structure(p.DbPlayer.LoginDays, &p.LoginDay)
}

// SaveData 缓存玩家数据到mysql
func (p *Player) SaveData() bool {
	p.DbPlayer.BackPack = util.ToString(p.BackPack)
	p.DbPlayer.CheckPointRecord = util.ToString(p.CheckPointRecords)
	p.DbPlayer.LoginDays = util.ToString(p.LoginDay)
	tx := db.DBC.Begin()
	err := tx.Save(p.DbPlayer).Where("id = ?", p.DbPlayer.ID).Error
	if err != nil {
		log.Release("save player error ", err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

// SevenDaysLogin 领取七日签到奖励
func (p *Player) SevenDaysLogin() {
	rewards := make(map[int32]int32)
	for k, v := range p.LoginDay {
		if v > 0 {
			index := common.GetCf(common.LOGIN_SEVEN_DAY).Index(k + 1)
			if index == nil {
				p.Agent.WriteMsg(&pb.SevenDaysRewardRsp{Code: int32(pb.Error_ReadConfigTable)})
				return
			}
			sevenDay := index.(*gamedata.Login_seven_day)
			for k1, v1 := range sevenDay.Rewardprops {
				rewards[int32(v1)] += int32(sevenDay.Rewardcounts[k1])
				p.BackPack[int32(v1)] += int32(sevenDay.Rewardcounts[k1])
			}
		}
	}
	p.Agent.WriteMsg(&pb.SevenDaysRewardRsp{
		Code:    int32(pb.Error_No),
		Rewards: rewards,
	})
}

// StarsRankList 玩家排行
func (p *Player) StarsRankList(num int32) {
	list := db.GetStarsRankList(int64(num))
	pbRankList := make([]*pb.RankStruct, 0)
	for _, v := range list {
		pbRankList = append(pbRankList, &pb.RankStruct{
			Uid:    v.Uid,
			Avatar: v.Avatar,
			Name:   v.Name,
			Stars:  int32(v.Stars),
		})
	}
	p.Agent.WriteMsg(&pb.StarsRankListRsp{
		Code:     int32(pb.Error_No),
		RankList: pbRankList,
	})
}

// SwitchPlate 切换挡板
func (p *Player) SwitchPlate(req *pb.SwitchBafflePlateReq) {
	if _, has := p.BackPack[req.NewId]; has && p.BackPack[req.NewId] != 0 {
		p.DbPlayer.CurrPlate = req.NewId
		p.Agent.WriteMsg(&pb.SwitchBafflePlateRsp{
			Code:        int32(pb.Error_No),
			CurrPlateId: p.DbPlayer.CurrPlate,
			Counts:      p.BackPack[req.NewId],
		})
	} else {
		p.Agent.WriteMsg(&pb.SwitchBafflePlateRsp{Code: int32(pb.Error_Req)})
	}
}

func (p *Player) GMEditor(cmd string) {
	gmTp := cmd[:5]
	switch gmTp {
	case "props":
		propId, _ := strconv.Atoi(cmd[5:11])
		counts, _ := strconv.Atoi(cmd[11:])
		p.BackPack[int32(propId)] += int32(counts)
		p.Agent.WriteMsg(&pb.NoticePropsChange{Backpack: p.BackPack})
	}
}

func (p *Player) ShopBuying(req *pb.ShopBuyingReq) {
	index := common.GetCf(common.SHOP).Index(int(req.PropId))
	if index == nil {
		p.Agent.WriteMsg(&pb.ShopBuyingRsp{Code: int32(pb.Error_ReadConfigTable)})
		return
	}
	shop := index.(*gamedata.Shop)
	for k, v := range shop.Costid {
		if p.BackPack[int32(v)] < int32(shop.Costnum[k]) {
			p.Agent.WriteMsg(&pb.ShopBuyingRsp{Code: int32(pb.Error_Req)})
			return
		}
	}
	for k, v := range shop.Costid {
		p.BackPack[int32(v)] -= int32(shop.Costnum[k])
	}
	rewards := make(map[int32]int32)
	for k, v := range shop.Proptype {
		rewards[int32(v)] += int32(shop.Propnum[k])
		p.BackPack[int32(v)] += int32(shop.Propnum[k])
	}
	p.Agent.WriteMsg(&pb.ShopBuyingRsp{
		Code:     int32(pb.Error_No),
		BackPack: p.BackPack,
		Rewards:  rewards,
	})
}

// SettleRank 结算排行榜
func SettleRank() {
	db.DBRedis.Del("rank")
	db.DBRedis.SUnionStore("rank", "stars")
}

// GetPersonalRank 获取结算名次
func (p *Player) GetPersonalRank() {
	list := db.GetYesterdayStarsRankList(-1)
	tag := -1
	for k, v := range list {
		if v.Uid == p.DbPlayer.Uid {
			tag = k
			break
		}
	}
	if tag >= 0 {
		tag = tag + 1
	}
	p.Agent.WriteMsg(&pb.YesterdayRankRsp{
		Code: int32(pb.Error_No),
		Rank: int32(tag),
	})
}

// YesterdayRankReward 领取昨日排行榜奖励
func (p *Player) YesterdayRankReward() {
	clock := util.GetZeroClock()
	if p.DbPlayer.RankReceive > int32(clock) && p.DbPlayer.RankReceive < int32(clock)+24*3600 {
		p.Agent.WriteMsg(&pb.ReceiveRankRewardRsp{Code: int32(pb.Error_Req)})
	}
	list := db.GetYesterdayStarsRankList(-1)
	tag := -1
	for k, v := range list {
		if v.Uid == p.DbPlayer.Uid {
			tag = k
			break
		}
	}
	if tag >= 0 {
		tag = tag + 1
	}
	//读表发奖励
	index := common.GetCf(common.RANK).Index(tag)
	if index == nil {
		p.Agent.WriteMsg(&pb.ReceiveRankRewardRsp{Code: int32(pb.Error_Req)})
		return
	}
	rank := index.(*gamedata.Rank)
	rewards := make(map[int32]int32)
	for _, v := range rank.Rewards {
		rewards[int32(v[0])] += int32(v[1])
		p.BackPack[int32(v[0])] += int32(v[1])
	}
	p.DbPlayer.RankReceive = int32(time.Now().Unix())
	p.Agent.WriteMsg(&pb.ReceiveRankRewardRsp{
		Code:    int32(pb.Error_No),
		Rewards: rewards,
		RankTag: p.DbPlayer.RankReceive,
	})
}

func (p *Player) WatchAdv(advId int) {
	shop := common.GetCf(common.SHOP).Index(advId).(*gamedata.Shop)
	rewards := make(map[int32]int32)
	for k, v := range shop.Proptype {
		rewards[int32(v)] += int32(shop.Propnum[k])
		p.BackPack[int32(v)] += int32(shop.Propnum[k])
	}
	p.Agent.WriteMsg(&pb.WatchAdvRewardsRsp{
		Code:     int32(pb.Error_No),
		Rewards:  rewards,
		Backpack: p.BackPack,
	})
}
