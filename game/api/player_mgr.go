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
	PMap  map[string]gate.Agent
	mutex sync.RWMutex
}

func NewPlayerManager() *PlayerManger {
	mgr := &PlayerManger{
		PMap:  make(map[string]gate.Agent),
		mutex: sync.RWMutex{},
	}
	return mgr
}

func (pMgr *PlayerManger) UserLogin(player *db.Player, a gate.Agent, req *pb.LoginReq) {
	newPlayer := NewPlayer(player, a)
	newPlayer.LoadData()
	a.SetUserData(newPlayer)
	pMgr.mutex.Lock()
	defer pMgr.mutex.Unlock()
	pMgr.PMap[player.Uid] = a
	pMgr.JudgeTodayLogin(newPlayer)
	db.UpdatePlayerBaseInfo(req.Uid, req.Avatar, req.Name)
	newPlayer.Agent.WriteMsg(&pb.LoginRsp{
		Code: int32(pb.Error_No),
		PlayerInfo: &pb.Player{
			Uid:              newPlayer.DbPlayer.Uid,
			BackPack:         newPlayer.BackPack,
			CheckPointRecord: newPlayer.CheckPointRecords,
			LoginDays:        newPlayer.LoginDay,
		},
	})
}

func (pMgr *PlayerManger) CreateRole(req *pb.LoginReq, a gate.Agent) {
	player := &db.Player{
		Uid:              req.Uid,
		BackPack:         util.ToString(make(map[int32]int32)),
		Gold:             0,
		CheckPointRecord: util.ToString(make(map[int32]int32)),
		LoginDays:        util.ToString(make([]int32, 0)),
	}
	err := db.DBC.Create(player).Error
	if err != nil {
		log.Release("创建角色失败", err)
		a.WriteMsg(&pb.LoginRsp{Code: int32(pb.Error_CreateRole)})
		return
	}
	newPlayer := NewPlayer(player, a)
	newPlayer.LoadData()
	db.UpdatePlayerBaseInfo(req.Uid, req.Avatar, req.Name)
	newPlayer.Agent.WriteMsg(&pb.LoginRsp{
		Code: int32(pb.Error_No),
		PlayerInfo: &pb.Player{
			Uid:              newPlayer.DbPlayer.Uid,
			BackPack:         newPlayer.BackPack,
			CheckPointRecord: newPlayer.CheckPointRecords,
			LoginDays:        newPlayer.LoginDay,
		},
	})
}

func (pMgr *PlayerManger) JudgeTodayLogin(p *Player) {
	today := util.GetZeroClock()
	tag := true
	for _, v := range p.LoginDay {
		if v == int32(today) || -v == int32(today) {
			tag = false
		}
	}
	if tag == true {
		p.LoginDay = append(p.LoginDay, int32(today))
	}
}

func (pMgr *PlayerManger) StarsRankList(p *Player, req *pb.StarsRankListReq) {
	p.StarsRankList(req.Num)
}

func (pMgr *PlayerManger) SevenDaysLogin(p *Player) {
	p.SevenDaysLogin()
}
