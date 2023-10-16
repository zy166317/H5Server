package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"leaf_server/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	Init()
}

func (m *Module) OnDestroy() {
	log.Release("closed")
}
