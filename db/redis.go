package db

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

type ChallengeRecord struct {
	Uid            string
	Lv             int32
	Name           string
	Success        bool
	StartTime      int64
	ProtectionTime int64
	GetScore       int64
	IsRevenge      int //1未被复仇成功 2 复仇成功
	Realm          int32
}

type challengeArr []ChallengeRecord

// Len 实现Interface接口，就是实现接口中的所有方法
func (ss challengeArr) Len() int {
	return len(ss)
}

// Less Less方法就是决定使用什么标准进行排序
// 按时间进行排序
func (ss challengeArr) Less(i, j int) bool {
	return ss[i].StartTime > ss[j].StartTime
}

func (ss challengeArr) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i] //与上面的交换方法一致

}

type RankInfo struct {
	Uid    string
	Avatar string
	Name   string
	Stars  int
}

func UpdatePlayerBaseInfo(uid, avatar, name string) {
	data := &RankInfo{
		Uid:    uid,
		Avatar: avatar,
		Name:   name,
		Stars:  0,
	}
	bytes, _ := json.Marshal(data)
	DBRedis.Set("uid", bytes, -1)
}

func GetStarsRankList(num int64) []*RankInfo {
	rankList := make([]*RankInfo, 0)
	val := DBRedis.ZRevRangeWithScores("stars", 0, num).Val()
	for _, v := range val {
		bytes, _ := DBRedis.Get(v.Member.(string)).Bytes()
		rankInfo := &RankInfo{}
		json.Unmarshal(bytes, rankInfo)
		rankInfo.Stars = int(v.Score)
		rankList = append(rankList, rankInfo)
	}
	return rankList
}

func UpdatePlayerStars(stars int, uid string) {
	rank := []redis.Z{
		{
			Score:  float64(stars),
			Member: uid,
		},
	}
	DBRedis.ZAdd("stars", rank...)
}

func GetYesterdayStarsRankList(num int64) []*RankInfo {
	rankList := make([]*RankInfo, 0)
	val := DBRedis.ZRevRangeWithScores("rank", 0, num).Val()
	for _, v := range val {
		bytes, _ := DBRedis.Get(v.Member.(string)).Bytes()
		rankInfo := &RankInfo{}
		json.Unmarshal(bytes, rankInfo)
		rankInfo.Stars = int(v.Score)
		rankList = append(rankList, rankInfo)
	}
	return rankList
}
