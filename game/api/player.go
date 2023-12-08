package api

import (
	"github.com/name5566/leaf/gate"
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

func (p *Player) LoadData() {
	util.ConvertString2Structure(p.DbPlayer.BallPack, &p.BallPack)
	util.ConvertString2Structure(p.DbPlayer.CheckPointRecord, &p.CheckPointRecords)
	util.ConvertString2Structure(p.DbPlayer.LoginDays, &p.LoginDay)
}
