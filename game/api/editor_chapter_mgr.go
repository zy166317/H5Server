package api

import (
	"github.com/name5566/leaf/gate"
	"leaf_server/proto/pb"
	"sync"
)

var EcMgr = NewEditorChapterManger()

type EditorChapter struct {
}

type EditorChapterManger struct {
	mutex sync.RWMutex
}

func NewEditorChapterManger() *EditorChapterManger {
	return &EditorChapterManger{mutex: sync.RWMutex{}}
}

func (ecMgr *EditorChapterManger) EditorChapterInfo(a gate.Agent, req *pb.EditorChapterReq) {
	if req.ChapterId > 0 {
		ChangeChapter(a, req)
	} else {
		EditorChapterAdd(a, req)
	}
}

func (ecMgr *EditorChapterManger) GetAllChapterInfo(a gate.Agent) {
	GetAllChapter(a, nil)
}

func (ecMgr *EditorChapterManger) ChapterDetail(a gate.Agent, req *pb.ChapterDetailByIdReq) {
	GetChapterById(a, req)
}
