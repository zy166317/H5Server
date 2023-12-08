package api

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leaf_server/db"
	"leaf_server/proto/pb"
	"leaf_server/util"
	"time"
)

// EditorChapterAdd 新增关卡
func EditorChapterAdd(a gate.Agent, req *pb.EditorChapterReq) {
	pos := make([][]int32, 0)
	color := make([]int32, 0)
	Tps := make([]int32, 0)
	for _, v := range req.Info {
		pos = append(pos, []int32{v.Row, v.Col})
		color = append(color, v.Color)
		Tps = append(Tps, v.Type)
	}
	err := db.DBC.Create(&db.Chapter{
		Pos:       util.ToString(pos),
		Color:     util.ToString(color),
		Types:     util.ToString(Tps),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}).Error
	if err != nil {
		log.Release("create ball error ", err)
	}
	a.WriteMsg(&pb.EditorChapterRsp{Code: int32(pb.Error_No)})
}

func ChangeChapter(a gate.Agent, req *pb.EditorChapterReq) {
	pos := make([][]int32, 0)
	color := make([]int32, 0)
	Tps := make([]int32, 0)
	for _, v := range req.Info {
		pos = append(pos, []int32{v.Row, v.Col})
		color = append(color, v.Color)
		Tps = append(Tps, v.Type)
	}
	chapter := &db.Chapter{
		Pos:   util.ToString(pos),
		Color: util.ToString(color),
		Types: util.ToString(Tps),
	}
	err := db.DBC.Model(&db.Chapter{}).Where("id = ?", req.ChapterId).Updates(chapter).Error
	if err != nil {
		log.Release("change chapter error", err)
		a.WriteMsg(&pb.EditorChapterRsp{Code: int32(pb.Error_Req)})
		return
	}
	a.WriteMsg(&pb.EditorChapterRsp{Code: int32(pb.Error_No)})
}

func GetAllChapter(a gate.Agent, req *pb.GetAllChapterReq) {
	chapters := make([]*db.Chapter, 0)
	err := db.DBC.Model(&db.Chapter{}).Where("id>0").Find(&chapters).Error
	if err != nil {
		log.Error("Find chapter error ", err)
		a.WriteMsg(&pb.GetAllChapterRsp{
			Code: int32(pb.Error_Mysql),
		})
		return
	}
	chapterIds := make([]int32, 0)
	for _, v := range chapters {
		chapterIds = append(chapterIds, int32(v.Id))

	}
	a.WriteMsg(&pb.GetAllChapterRsp{
		Code:       int32(pb.Error_No),
		ChapterIds: chapterIds,
	})
}

func GetChapterById(a gate.Agent, req *pb.ChapterDetailByIdReq) {
	var chapterInfo *db.Chapter
	err := db.DBC.Model(&db.Chapter{}).Where("id = ?", req.ChapterId).First(&chapterInfo).Error
	if err != nil {
		log.Error("Find chapter error ", err)
		a.WriteMsg(&pb.ChapterDetailByIdRsp{
			Code: int32(pb.Error_Mysql),
		})
		return
	}
	color := make([]int32, 0)
	pos := make([][]int32, 0)
	Tps := make([]int32, 0)
	util.ConvertString2Structure(chapterInfo.Color, &color)
	util.ConvertString2Structure(chapterInfo.Pos, &pos)
	util.ConvertString2Structure(chapterInfo.Types, &Tps)
	pbInfo := make([]*pb.Chapter, 0)
	for i := 0; i < len(color); i++ {
		pbInfo = append(pbInfo, &pb.Chapter{
			Color: color[i],
			Row:   pos[i][0],
			Col:   pos[i][1],
			Type:  Tps[i],
		})
	}
	a.WriteMsg(&pb.ChapterDetailByIdRsp{
		Code:      int32(pb.Error_No),
		Info:      &pb.ChapterDetail{Info: pbInfo},
		ChapterId: req.ChapterId,
	})
}
