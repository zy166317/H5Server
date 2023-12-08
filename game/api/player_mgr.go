package api

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leaf_server/db"
	"leaf_server/proto/pb"
	"leaf_server/util"
	"sync"
)

var (
	PlayerMgr = NewPlayerManager()
)

type PlayerManger struct {
	mutex sync.RWMutex
}

func NewPlayerManager() *PlayerManger {
	mgr := &PlayerManger{
		mutex: sync.RWMutex{},
	}
	return mgr
}

func (pMgr *PlayerManger) UserLogin(player *db.Player, a gate.Agent) {
	newPlayer := NewPlayer(player, a)
	newPlayer.LoadData()
	newPlayer.Agent.WriteMsg(&pb.LoginRsp{
		Code: int32(pb.Error_No),
		PlayerInfo: &pb.Player{
			Uid:              newPlayer.DbPlayer.Uid,
			BallPack:         newPlayer.BallPack,
			CheckPointRecord: newPlayer.CheckPointRecords,
			LoginDays:        newPlayer.LoginDay,
		},
	})
}

func (pMgr *PlayerManger) CreateRole(uid string, a gate.Agent) {
	player := &db.Player{
		Uid:              uid,
		BallPack:         util.ToString(make(map[int32]int32)),
		Gold:             0,
		CheckPointRecord: util.ToString(make(map[int32]int32)),
		LoginDays:        util.ToString(make(map[int32]bool)),
	}
	err := db.DBC.Create(player).Error
	if err != nil {
		log.Release("创建角色失败", err)
		a.WriteMsg(&pb.LoginRsp{Code: int32(pb.Error_CreateRole)})
		return
	}
	newPlayer := NewPlayer(player, a)
	newPlayer.LoadData()
	newPlayer.Agent.WriteMsg(&pb.LoginRsp{
		Code: int32(pb.Error_No),
		PlayerInfo: &pb.Player{
			Uid:              newPlayer.DbPlayer.Uid,
			BallPack:         newPlayer.BallPack,
			CheckPointRecord: newPlayer.CheckPointRecords,
			LoginDays:        newPlayer.LoginDay,
		},
	})
}
