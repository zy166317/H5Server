package util

//
//import (
//	"fmt"
//	"leaf_server/common"
//	"leaf_server/gamedata"
//	rand "math/rand"
//	"strconv"
//	"time"
//)
//
//type RandomIntArr struct {
//	Produce int
//	Weight  int
//}
//
//type IndexRandom struct {
//	Produce int
//	Weight  int
//	Index   int
//}
//
//func RandomTimesByQuality(quality int) int {
//	parmas := make(map[int]int)
//	for i := 1; i < 10; i++ {
//		parmas[i] = i - 1
//	}
//	return parmas[quality]
//}
//
//// RandomEquip
//// 随机开鼎产出装备类型
//// 返回值为装备类型type
//func RandomEquip() int {
//	arr := make([]RandomIntArr, 0)
//	for i := 1; i <= 12; i++ {
//		arr = append(arr, RandomIntArr{
//			Produce: i,
//			Weight:  10,
//		})
//	}
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedEquip RandomIntArr
//	for _, equip := range arr {
//		if r < equip.Weight {
//			selectedEquip = equip
//			break
//		}
//		r -= equip.Weight
//	}
//	return selectedEquip.Produce
//}
//
//// RandomEquipLevel
//// 开鼎产出装备等级随机计算
//// 返回值为计算后的level
//func RandomEquipLevel(level int) int {
//	arr := make([]RandomIntArr, 0)
//	arr = append(arr, RandomIntArr{
//		Produce: level - 1,
//		Weight:  25,
//	},
//		RandomIntArr{
//			Produce: level,
//			Weight:  100,
//		},
//		RandomIntArr{
//			Produce: level + 1,
//			Weight:  15,
//		},
//		RandomIntArr{
//			Produce: level + 2,
//			Weight:  5,
//		},
//	)
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedEquip RandomIntArr
//	for _, equip := range arr {
//		if r < equip.Weight {
//			selectedEquip = equip
//			break
//		}
//		r -= equip.Weight
//	}
//
//	return selectedEquip.Produce
//}
//
//// RandomQuality
//// 品质随机计算
//// 返回值品质id
//func RandomQuality(level int) int {
//
//	arr := make([]RandomIntArr, 0)
//	tripodBox := common.GetCf(common.TRIPOD_BOX).Index(level).(*gamedata.TripodBox)
//	for i, v := range tripodBox.Weightshow {
//		arr = append(arr, RandomIntArr{
//			Produce: i + 1,
//			Weight:  int(v * float32(100)),
//		})
//	}
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedEquip RandomIntArr
//	for _, equip := range arr {
//		if r < equip.Weight {
//			selectedEquip = equip
//			break
//		}
//		r -= equip.Weight
//	}
//
//	return selectedEquip.Produce
//}
//
//// RandomFirstAttr 第一属性随机计算
//// 返回值为计算后的第一属性固定值切片
//func RandomFirstAttr(level, quality int) []int {
//	random := common.GetCf(common.EQUIP_RANDOM).Index(level).(*gamedata.EquipRandom)
//	arr := make([]int, 4)
//	base := make([]float32, 4)
//	for i, v := range random.Random {
//		a := rand.Intn(int(v * 1000))
//		base[i] = random.Base[i] + float32(a/1000)
//	}
//
//	times := RandomTimesByQuality(quality)
//	if times > 0 {
//		for i, k := range random.Quality {
//			base[i] += k * float32(times)
//			arr[i] = int(base[i])
//		}
//	} else {
//		for i, v := range base {
//			arr[i] = int(v)
//		}
//	}
//	return arr
//}
//
//// RandomSecondAttr 第二属性随机计算
//// 返回值为属性id：值
//func RandomSecondAttr(level, quality int) (int, float32) {
//	if quality <= 1 {
//		return 0, 0
//	}
//	arr := make([]RandomIntArr, 0)
//	random := common.GetCf(common.EQUIP_RANDOM).Index(level).(*gamedata.EquipRandom)
//	for i, v := range random.Secondweight {
//		arr = append(arr, RandomIntArr{
//			Produce: i + 5,
//			Weight:  v,
//		})
//	}
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedEquip RandomIntArr
//	for _, equip := range arr {
//		if r < equip.Weight {
//			selectedEquip = equip
//			break
//		}
//		r -= equip.Weight
//	}
//	secondMapTimes := make(map[int]int)
//	secondMapTimes[1] = 0
//	secondMapTimes[2] = 0
//	secondMapTimes[3] = 1
//	secondMapTimes[4] = 2
//	secondMapTimes[5] = 3
//	secondMapTimes[6] = 4
//	secondMapTimes[7] = 5
//	secondMapTimes[8] = 6
//	secondMapTimes[9] = 8
//	t := secondMapTimes[quality]
//	//第二属性随机上浮值计算
//	baseAttr := random.Secondbase
//	if t > 0 {
//		intn := rand.Intn(50)
//		secRand := float32(intn) / float32(100)
//		baseAttr = random.Secondbase * (secRand + 1)
//		for i := 0; i < t; i++ {
//			baseAttr = baseAttr * random.Secondgrow
//		}
//	}
//	var num float64
//	if baseAttr > 0.01 {
//		num, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", baseAttr), 64)
//	} else {
//		num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", baseAttr), 64)
//	}
//	return selectedEquip.Produce, float32(num)
//}
//
//// RandomThirdAttr
//// 第三属性随机计算
//// 返回值为属性id和属性值
//func RandomThirdAttr(level, Quality int) (int, float32) {
//	if Quality < 5 {
//		return 0, 0
//	}
//	arr := make([]RandomIntArr, 0)
//	random := common.GetCf(common.EQUIP_RANDOM).Index(level).(*gamedata.EquipRandom)
//	for i, v := range random.Thirdweight {
//		arr = append(arr, RandomIntArr{
//			Produce: i + 301,
//			Weight:  v,
//		})
//	}
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedEquip RandomIntArr
//	for _, equip := range arr {
//		if r < equip.Weight {
//			selectedEquip = equip
//			break
//		}
//		r -= equip.Weight
//	}
//	thirdMap := make(map[int]int)
//	thirdMap[1] = 0
//	thirdMap[2] = 0
//	thirdMap[3] = 0
//	thirdMap[4] = 0
//	thirdMap[5] = 0
//	thirdMap[6] = 1
//	thirdMap[7] = 2
//	thirdMap[8] = 3
//	thirdMap[9] = 4
//	thirdBase := random.Thirdbase
//	times := thirdMap[Quality]
//	if times > 0 {
//		thirdInt := rand.Intn(100)
//		thirdRand := float32(thirdInt) / float32(100)
//		thirdBase = random.Thirdbase * (thirdRand + 1)
//		for i := 0; i < times; i++ {
//			thirdBase = thirdBase * random.Thirdgrow
//		}
//	}
//	var num float64
//	if thirdBase > 0.01 {
//		num, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", thirdBase), 64)
//	} else {
//		num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", thirdBase), 64)
//	}
//
//	return selectedEquip.Produce, float32(num)
//}
//
//func RandomBreakThrough(sPercent, fPercent int) int {
//	arr := make([]RandomIntArr, 0)
//	arr = append(arr, RandomIntArr{
//		Produce: 1,
//		Weight:  sPercent,
//	},
//		RandomIntArr{
//			Produce: 0,
//			Weight:  fPercent,
//		},
//	)
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedPercent RandomIntArr
//	for _, percent := range arr {
//		if r < percent.Weight {
//			selectedPercent = percent
//			break
//		}
//		r -= percent.Weight
//	}
//
//	return selectedPercent.Produce
//}
//
//func IsOrNotSuccess(sPercent, fPercent int) int {
//	arr := make([]RandomIntArr, 0)
//	arr = append(arr, RandomIntArr{
//		Produce: 1,
//		Weight:  sPercent,
//	},
//		RandomIntArr{
//			Produce: 0,
//			Weight:  fPercent,
//		},
//	)
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedPercent RandomIntArr
//	for _, percent := range arr {
//		if r < percent.Weight {
//			selectedPercent = percent
//			break
//		}
//		r -= percent.Weight
//	}
//
//	return selectedPercent.Produce
//}
//
//// RandomGrassId
//// 随机生成草药
//// 返回值为草药id
//func RandomGrassId() int {
//	arr := make([]RandomIntArr, 0)
//	arr = append(arr, RandomIntArr{
//		Produce: 51001,
//		Weight:  60,
//	},
//		RandomIntArr{
//			Produce: 51001,
//			Weight:  60,
//		},
//	)
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedEquip RandomIntArr
//	for _, equip := range arr {
//		if r < equip.Weight {
//			selectedEquip = equip
//			break
//		}
//		r -= equip.Weight
//	}
//
//	return selectedEquip.Produce
//}
//
//func RandomGrassIdByInterval(firId, secId int) int {
//	arr := make([]RandomIntArr, 0)
//	arr = append(arr, RandomIntArr{
//		Produce: firId,
//		Weight:  50,
//	},
//		RandomIntArr{
//			Produce: secId,
//			Weight:  50,
//		},
//	)
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedPercent RandomIntArr
//	for _, percent := range arr {
//		if r < percent.Weight {
//			selectedPercent = percent
//			break
//		}
//		r -= percent.Weight
//	}
//
//	return selectedPercent.Produce
//}
//
//// RandomTalismanExp 本命法宝升级经验值随机
//func RandomTalismanExp() int {
//	arr := make([]RandomIntArr, 0)
//	arr = append(arr, RandomIntArr{
//		Produce: 1,
//		Weight:  60,
//	},
//		RandomIntArr{
//			Produce: 2,
//			Weight:  20,
//		},
//		RandomIntArr{
//			Produce: 3,
//			Weight:  10,
//		},
//		RandomIntArr{
//			Produce: 5,
//			Weight:  5,
//		},
//		RandomIntArr{
//			Produce: 10,
//			Weight:  5,
//		},
//	)
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedEquip RandomIntArr
//	for _, equip := range arr {
//		if r < equip.Weight {
//			selectedEquip = equip
//			break
//		}
//		r -= equip.Weight
//	}
//
//	return selectedEquip.Produce
//}
//
//// RandomOrderByDay 符箓每日订单任务
//func RandomOrderByDay() map[int32]int32 {
//	arr := make([]IndexRandom, 0)
//	cf := common.GetCf(common.TALLY_ORDER_TASK)
//	for i := 1; i < cf.NumRecord(); i++ {
//		task := cf.Record(i).(*gamedata.TallyOrderTask)
//		arr = append(arr, IndexRandom{
//			Produce: task.Id,
//			Weight:  task.Weight,
//			Index:   i,
//		})
//	}
//	firstData := Random(arr)
//	var secArr []IndexRandom
//	secArr = arr[:firstData.Index]
//	secArr = append(secArr, arr[firstData.Index+1:]...)
//	secData := Random(secArr)
//	thirdArr := secArr[:secData.Index]
//	thirdArr = append(thirdArr, secArr[secData.Index+1:]...)
//	thirdData := Random(thirdArr)
//	result := make(map[int32]int32)
//	result[int32(firstData.Produce)] = 0
//	result[int32(secData.Produce)] = 0
//	result[int32(thirdData.Produce)] = 0
//	if len(result) != 3 {
//		RandomOrderByDay()
//	}
//
//	return result
//}
//
//func RandomOneOrder(currOrder map[int32]int32) int {
//	arr := make([]IndexRandom, 0)
//	cf := common.GetCf(common.TALLY_ORDER_TASK)
//	for i := 0; i < cf.NumRecord(); i++ {
//		task := cf.Record(i).(*gamedata.TallyOrderTask)
//		_, has := currOrder[int32(task.Id)]
//		if !has {
//			arr = append(arr, IndexRandom{
//				Produce: task.Id,
//				Weight:  task.Weight,
//				Index:   i,
//			})
//		}
//	}
//	random := Random(arr)
//	return random.Produce
//}
//
//func Random(arr []IndexRandom) IndexRandom {
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedEquip IndexRandom
//	for _, equip := range arr {
//		if r < equip.Weight {
//			selectedEquip = equip
//			break
//		}
//		r -= equip.Weight
//	}
//
//	return selectedEquip
//}
//
//func RandomDoMakeFuLu(scene *gamedata.TallyScene, doFuLuAttr map[int32]float32, currPaper int32) int {
//	arrFuLu := make([]RandomIntArr, 0)
//	weight := 0
//	for i, v := range scene.Fishes {
//		tally := common.GetCf(common.ITEM).Index(v).(*gamedata.Item)
//		switch tally.Subtype {
//		case 1:
//			weight = int(float32(scene.Fishweight[i]) * (1 + doFuLuAttr[608]) * (1 - doFuLuAttr[614]))
//		case 2:
//			weight = int(float32(scene.Fishweight[i]) * (1 + doFuLuAttr[609]) * (1 - doFuLuAttr[615]))
//		case 3:
//			weight = int(float32(scene.Fishweight[i]) * (1 + doFuLuAttr[610]) * (1 - doFuLuAttr[616]))
//		case 4:
//			weight = int(float32(scene.Fishweight[i]) * (1 + doFuLuAttr[611]))
//		case 5:
//			weight = int(float32(scene.Fishweight[i]) * (1 + doFuLuAttr[613]) * (1 - doFuLuAttr[604]))
//		case 6:
//			weight = scene.Fishweight[i]
//		}
//		arrFuLu = append(arrFuLu, RandomIntArr{
//			Produce: v,
//			Weight:  weight,
//		})
//	}
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arrFuLu {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedEquip RandomIntArr
//	for _, equip := range arrFuLu {
//		if r < equip.Weight {
//			selectedEquip = equip
//			break
//		}
//		r -= equip.Weight
//	}
//	tally := common.GetCf(common.TALLY).Index(selectedEquip.Produce).(*gamedata.Tally)
//	if tally.Needbaits > 0 && int32(tally.Needbaits) != currPaper {
//		RandomDoMakeFuLu(scene, doFuLuAttr, currPaper)
//	}
//	return selectedEquip.Produce
//}
//
//func RandomGetPvpPlayer(rankIndex int) []int {
//	indexArr := make([]int, 0)
//	cf := common.GetCf(common.ARENA_MATE)
//	for i := 0; i < cf.NumRecord(); i++ {
//		info := cf.Record(i).(*gamedata.ArenaMate)
//		if rankIndex <= info.Ranktop && rankIndex >= info.Rankbottom {
//			for _, v := range info.Rankpart {
//				rank := getConfigRank(rankIndex, v)
//				indexArr = append(indexArr, rank)
//			}
//		}
//	}
//	return indexArr
//}
//
//// 排行榜区间确定排名
//func getConfigRank(baseRand int, v []int) int {
//	tempRank := 0 //临时的排名
//	if v[0] == v[1] {
//		tempRank = v[0]
//	}
//	if v[0] > v[1] {
//		tempRank = v[1] + rand.Intn(v[0]-v[1]+1)
//	} else if v[0] < v[1] {
//		tempRank = v[0] + rand.Intn(v[1]-v[0]+1)
//	} else {
//		tempRank = v[0]
//	}
//	return baseRand + tempRank
//}
//
//// OpenTripodBoxSelectId 每日开鼎中将随机
//func OpenTripodBoxSelectId() []int32 {
//	config := common.GetCf(common.GLOBAL_CONFIG).Index("ArenaInitialBoxLoot").(*gamedata.Globalconfig)
//	selectId := make([]int32, 0)
//	for _, v := range config.Data3 {
//		num := RandInt(v[0], v[1])
//		selectId = append(selectId, int32(num))
//	}
//	return selectId
//}
//
//// RandInt 生成指定区间随机数
//func RandInt(min, max int) int {
//	if min >= max || min == 0 || max == 0 {
//		return max
//	}
//	return rand.Intn(max-min+1) + min
//}
//
//// RandKfAttr 功法升阶随机加成
//func RandKfAttr(kfId int) (int, int) {
//	kfConfig := common.GetCf(common.KUNGFU).Index(kfId).(*gamedata.KungFu)
//	arr := make([]IndexRandom, 0)
//	for i, v := range kfConfig.Randomattributetype {
//		arr = append(arr, IndexRandom{
//			Produce: v,
//			Weight:  kfConfig.Randomattributeweight[i],
//			Index:   i,
//		})
//	}
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedEquip IndexRandom
//	for _, equip := range arr {
//		if r < equip.Weight {
//			selectedEquip = equip
//			break
//		}
//		r -= equip.Weight
//	}
//
//	return selectedEquip.Produce, selectedEquip.Index
//}
//
//// RandomKfExplore 功法探索随机
//func RandomKfExplore(funcId, funcWeight string) []int32 {
//	globalconfigId := common.GetCf(common.GLOBAL_CONFIG).Index(funcId).(*gamedata.Globalconfig)
//	globalconfigWeight := common.GetCf(common.GLOBAL_CONFIG).Index(funcWeight).(*gamedata.Globalconfig)
//	kfArr := make([]IndexRandom, 0)
//	for k, v := range globalconfigId.Data2 {
//		kfArr = append(kfArr, IndexRandom{
//			Produce: v,
//			Weight:  globalconfigWeight.Data2[k],
//			Index:   k,
//		})
//	}
//	firstData := Random(kfArr)
//	var secArr []IndexRandom
//	secArr = kfArr[:firstData.Index]
//	secArr = append(secArr, kfArr[firstData.Index+1:]...)
//	secData := Random(secArr)
//	thirdArr := secArr[:secData.Index]
//	thirdArr = append(thirdArr, secArr[secData.Index+1:]...)
//	thirdData := Random(thirdArr)
//	var result []int32
//	result = append(result, int32(firstData.Produce), int32(secData.Produce), int32(thirdData.Produce))
//	return result
//}
//
//func RandomProgressIsFull(funcId, funcWeight string) []int32 {
//	globalconfigId := common.GetCf(common.GLOBAL_CONFIG).Index(funcId).(*gamedata.Globalconfig)
//	globalconfigWeight := common.GetCf(common.GLOBAL_CONFIG).Index(funcWeight).(*gamedata.Globalconfig)
//	kfArr := make([]IndexRandom, 0)
//	for k, v := range globalconfigId.Data2 {
//		kfArr = append(kfArr, IndexRandom{
//			Produce: v,
//			Weight:  globalconfigWeight.Data2[k],
//			Index:   k,
//		})
//	}
//	firstData := Random(kfArr)
//	var secArr []IndexRandom
//	secArr = kfArr[:firstData.Index]
//	secArr = append(secArr, kfArr[firstData.Index+1:]...)
//	secData := Random(secArr)
//	thirdArr := secArr[:secData.Index]
//	thirdArr = append(thirdArr, secArr[secData.Index+1:]...)
//	thirdData := Random(thirdArr)
//	var result []int32
//	result = append(result, int32(firstData.Produce), int32(secData.Produce), int32(thirdData.Produce))
//	return result
//}
//
//// RandomKfUnlock 初始化随机一个功法给玩家
//func RandomKfUnlock() int {
//	arr := make([]RandomIntArr, 0)
//	config := common.GetCf(common.GLOBAL_CONFIG).Index("kungFuFirstRefreshItem").(*gamedata.Globalconfig)
//	for _, v := range config.Data2 {
//		arr = append(arr, RandomIntArr{
//			Produce: v,
//			Weight:  1,
//		})
//	}
//	rand.Seed(time.Now().UnixNano())
//	totalWeight := 0
//	for _, prize := range arr {
//		totalWeight += prize.Weight
//	}
//
//	r := rand.Intn(totalWeight)
//	var selectedPercent RandomIntArr
//	for _, percent := range arr {
//		if r < percent.Weight {
//			selectedPercent = percent
//			break
//		}
//		r -= percent.Weight
//	}
//
//	return selectedPercent.Produce
//}
//
//// RandomFlEquipWash 符箓装备洗炼
//func RandomFlEquipWash(equipId, order int) map[int32]float32 {
//	randomCounts := 0
//	equip := common.GetCf(common.TALLY_EQUIP).Index(equipId).(*gamedata.TallyEquip)
//	for _, v := range equip.Recoinslotrate {
//		success := IsOrNotSuccess(v, 100-v)
//		if success == 1 {
//			randomCounts++
//		}
//	}
//	arr := make([]RandomIntArr, 0)
//	randomdata := make(map[int32]float32)
//	for _, v := range equip.Recoineffect {
//		arr = append(arr, RandomIntArr{
//			Produce: v,
//			Weight:  1,
//		})
//	}
//	for i := 0; i < randomCounts; i++ {
//		rand.Seed(time.Now().UnixNano())
//		totalWeight := 0
//		for _, prize := range arr {
//			totalWeight += prize.Weight
//		}
//
//		r := rand.Intn(totalWeight)
//		var selectedPercent RandomIntArr
//		for _, percent := range arr {
//			if r < percent.Weight {
//				selectedPercent = percent
//				break
//			}
//			r -= percent.Weight
//		}
//		rNum := rand.Intn(equip.Recoineffectqualitymax[order]-equip.Recoineffectqualitymin[order]+1) + equip.Recoineffectqualitymin[order]
//		if randomdata[int32(selectedPercent.Produce)] > 0 {
//			randomdata[int32(selectedPercent.Produce)] += float32(rNum) * equip.Recoineffectqualitybase
//		} else {
//			randomdata[int32(selectedPercent.Produce)] = float32(rNum) * equip.Recoineffectqualitybase
//		}
//	}
//	return randomdata
//}
