package api

import (
	"leaf_server/proto/pb"
	"sync"
)

var (
	CMgr = NewChallengeManager()
)

type ChallengeManger struct {
	mutex sync.RWMutex
}

func NewChallengeManager() *ChallengeManger {
	mgr := &ChallengeManger{
		mutex: sync.RWMutex{},
	}
	return mgr
}

func (cMgr *ChallengeManger) ChallengeChapter(p *Player, req *pb.ChallengeChapterReq) {
	p.ChallengeInfo.ChallengeChapter(p, req)
}

func (cMgr *ChallengeManger) UseProps(p *Player, req *pb.UsePropsReq) {
	p.ChallengeInfo.UseProps(p, req)
}
