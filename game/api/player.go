package api

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leaf_server/common"
	"leaf_server/db"
	"leaf_server/gamedata"
	"leaf_server/proto/pb"
	"leaf_server/util"
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
				rewards[int32(v1)] += int32(sevenDay.Rewardcount[k1])
				p.AddProps(int32(v1), int32(sevenDay.Rewardcount[k1]))
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

func (p *Player) AddProps(propId, propCounts int32) {
	if propId == 100000 {
		p.DbPlayer.Gold += propCounts
	} else {
		p.BackPack[propId] += propCounts
	}
}
