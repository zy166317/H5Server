package api

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leaf_server/db"
	"leaf_server/util"
)

type Player struct {
	Agent             gate.Agent
	DbPlayer          *db.Player
	ChallengeInfo     Challenge
	BallPack          map[int32]int32
	CheckPointRecords map[int32]int32
	LoginDay          map[int32]bool
}

func NewPlayer(player *db.Player, a gate.Agent) *Player {
	return &Player{
		Agent:    a,
		DbPlayer: player,
	}
}

// LoadData 加载玩家数据库数据
func (p *Player) LoadData() {
	util.ConvertString2Structure(p.DbPlayer.BallPack, &p.BallPack)
	util.ConvertString2Structure(p.DbPlayer.CheckPointRecord, &p.CheckPointRecords)
	util.ConvertString2Structure(p.DbPlayer.LoginDays, &p.LoginDay)
}

// SaveData 缓存玩家数据到mysql
func (p *Player) SaveData() bool {
	p.DbPlayer.BallPack = util.ToString(p.BallPack)
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
