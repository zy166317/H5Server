package internal

import (
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

	//连接mysql
	Init()

}

func (m *Module) OnDestroy() {
	//mgodb.Close()

}
