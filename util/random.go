package util

import (
	"math/rand"
	"sync"
	"time"
)

type IndexRandom struct {
	Produce int
	Weight  int
	Index   int
}

// Random 简单随机
func Random(arr []IndexRandom) IndexRandom {
	if len(arr) == 0 {
		return IndexRandom{}
	}
	rand.Seed(time.Now().UnixNano())
	totalWeight := 0
	for _, prize := range arr {
		totalWeight += prize.Weight
	}

	r := rand.Intn(totalWeight)
	var selectedEquip IndexRandom
	for _, equip := range arr {
		if r < equip.Weight {
			selectedEquip = equip
			break
		}
		r -= equip.Weight
	}
	return selectedEquip
}

var (
	//第一名称
	familyNames = make([]string, 0)

	//第二名称
	middleNamesArr = make([]string, 0)

	//定义一堆名字
	lastNames = make([]string, 0)

	//头像库
	AvatarLists = make([]string, 0)
)

func InitNameList() {
	//record := common.GetCf(common.RANDOM_NAME)
	//for i := 0; i < record.NumRecord(); i++ {
	//	name := record.Record(i).(*gamedata.Random_name)
	//	familyNames = append(familyNames, name.Firstname)
	//	middleNamesArr = append(middleNamesArr, name.Secondname)
	//	lastNames = append(lastNames, name.Thirdname)
	//}
	//cf := common.GetCf(common.HEAD_SCULPTURE)
	//for i := 0; i < cf.NumRecord(); i++ {
	//	head := cf.Record(i).(*gamedata.Head_Sculpture)
	//	if head.Type == 1 {
	//		AvatarLists = append(AvatarLists, strconv.Itoa(head.Id))
	//	}
	//}
}

func RandomUniqueName() string {
	name := GetRandomName()
	return name
}

func GetRandomName() (name string) {
	familyName := familyNames[GetRandomInt(0, len(familyNames)-1)]
	middleName := middleNamesArr[GetRandomInt(0, len(middleNamesArr)-1)]
	lastName := lastNames[GetRandomInt(0, len(lastNames)-1)]
	return familyName + middleName + lastName
}

var (
	//随机数互斥锁（确保GetRandomInt不能被并发访问）
	randomMutex sync.Mutex
)

// GetRandomInt /*获取[start,end]之间的随机数*/
func GetRandomInt(start, end int) int {
	//访问加同步锁，是因为并发访问时容易因为时间种子相同而生成相同的随机数，那就狠不随机鸟！
	randomMutex.Lock()

	//利用定时器阻塞1纳秒，保证时间种子得以更改
	<-time.After(1 * time.Nanosecond)

	//根据时间纳秒（种子）生成随机数对象
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//得到[start,end]之间的随机数
	n := start + r.Intn(end-start+1)

	//释放同步锁，供其它协程调用
	randomMutex.Unlock()
	return n
}

// GetRandomAvatar 随机获得一个头像
func GetRandomAvatar() string {
	randomInt := GetRandomInt(0, len(AvatarLists)-1)
	return AvatarLists[randomInt]
}
