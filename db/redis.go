package db

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
