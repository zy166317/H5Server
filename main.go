package main

import (
	lconf "github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"leaf_server/chat"
	"leaf_server/common"
	"leaf_server/conf"
	"leaf_server/db"
	"leaf_server/game"
	"leaf_server/gate"
	"leaf_server/http"
	"leaf_server/leaf"
	"leaf_server/login"
	http2 "net/http"
)

const defaultConfigPath = "./conf/conf.yml"

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath
	common.ConfigInit()
	common.ReloadAll()
	f, err := ioutil.ReadFile(defaultConfigPath)
	if err != nil {
		log.Release("read config file faild,path=%s, %s", defaultConfigPath, err)
		return
	}
	configuration := db.NewConfiguration()
	err = yaml.Unmarshal(f, configuration)
	if err != nil {
		log.Release("read config file faild,path=%s, %s", defaultConfigPath, err)
		return
	}
	err = db.Init(&configuration.DBGame)
	if err != nil {
		log.Release("db init  faild,path=%s", err)
		return
	}
	err = db.RedisInit(&configuration.DBRedis)
	if err != nil {
		log.Release("redis init failed,path=%s", err)
		return
	}
	RunGlobalFunction()
	go http.InitHttpService("127.0.0.1:9999")

	//pprof
	go func() {
		http2.ListenAndServe("localhost:6060", nil)
	}()
	//开启全局定时任务
	leaf.Run(
		gate.Module,
		game.Module,
		login.Module,
		chat.Module,
	)
}

// RunGlobalFunction 加载启动全局业务
func RunGlobalFunction() {
}

//func CalcResult(configId int) {
//	firGrids := [5][6]int32{}
//	secGrids := [5][6]int32{}
//	thiGrids := [5][6]int32{}
//	var firColor, secColor, thiColor int
//	index := common.GetCf(common.PlayResult1).Index(configId)
//	if index == nil {
//		return
//	}
//	result1 := index.(*gamedata.PlayResult1)
//	//判断是否中奖
//	if len(result1.Resultnum1) > 0 {
//		section := result1.Resultnum1[1] - result1.Resultnum1[0] + 1
//		//随机中奖花色个数
//		num := rand.Intn(section) + result1.Resultnum1[0]
//		//随机中奖花色
//		firColor = result1.Resultsuit1[rand.Intn(len(result1.Resultsuit1))]
//		firGrids = FirstAgileFillInGrid(firGrids, firColor, num)
//	}
//	//二次中奖
//	if len(result1.Resultnum2) > 0 {
//		section := result1.Resultnum2[1] - result1.Resultnum2[0] + 1
//		//随机中奖花色个数
//		num := rand.Intn(section) + result1.Resultnum2[0]
//		//随机中奖花色
//		secColor = result1.Resultsuit2[rand.Intn(len(result1.Resultsuit2))]
//		secGrids = SecondAgileFillInGrid(firGrids, secColor, firColor, num)
//	}
//	//三次中奖
//	if len(result1.Resultnum3) > 0 {
//		section := result1.Resultnum3[1] - result1.Resultnum3[0] + 1
//		//随机中奖花色个数
//		num := rand.Intn(section) + result1.Resultnum3[0]
//		//随机中奖花色
//		thiColor = result1.Resultsuit3[rand.Intn(len(result1.Resultsuit3))]
//		thiGrids = SecondAgileFillInGrid(secGrids, thiColor, secColor, num)
//	}
//	fmt.Println(firGrids)
//	fmt.Println(secGrids)
//	fmt.Println(thiGrids)
//}
//
//// FirstAgileFillInGrid 填充格子
//func FirstAgileFillInGrid(grids [5][6]int32, color, num int) [5][6]int32 {
//	var x, j int
//	//依据所有花色,随机生成30个不中奖的格子
//	gridsMap := make(map[int]int)
//	for i := 1; i < 31; i++ {
//		colorId := GetColor(0)
//		if gridsMap[colorId] <= 7 {
//			grids[x][j] = int32(colorId)
//			j++
//			if j >= 6 {
//				j = 0
//				x++
//			}
//			gridsMap[colorId] += 1
//		} else {
//			grids[x][j] = int32(colorId)
//			j++
//			if j >= 6 {
//				j = 0
//				x++
//			}
//			colorId = GetColor(colorId)
//			gridsMap[colorId] += 1
//		}
//	}
//	addNum := num - gridsMap[color]
//	toMap := DoubleArrToMap(grids)
//	for k, v := range toMap {
//		if addNum == 0 {
//			break
//		}
//		if v != int32(color) {
//			toMap[k] = int32(color)
//			addNum--
//		}
//	}
//	arr := MapToDoubleArr(toMap)
//	return arr
//}
//
//// SecondAgileFillInGrid 二次中奖格子填充
//func SecondAgileFillInGrid(grids [5][6]int32, color, oldColor, num int) [5][6]int32 {
//	colorNum := make(map[int32]int32)
//	for k, v := range grids {
//		for k1, v1 := range v {
//			if v1 == int32(oldColor) {
//				grids[k][k1] = 0
//			}
//			colorNum[v1] += 1
//		}
//	}
//	//删除第一次中奖花色之后的重新排列
//	newGrids := DelColor(grids)
//	//第二次中奖个数-当前已经拥有花色个数，其余的格子则使用数量最少的花色填充
//	newColor := int32(num) - colorNum[int32(color)]
//	colorArr := make([]int, 0)
//	//将中奖花色填入
//	for i := 0; i < int(newColor); i++ {
//		colorArr = append(colorArr, color)
//	}
//	//计算中奖之外的花色个数
//	subNum := num - int(newColor)
//	//填充剩下花色，保证都不中奖
//	delete(colorNum, int32(color))
//	for k, v := range colorNum {
//		if subNum <= 0 {
//			break
//		}
//		if v < 7 && subNum >= 7-int(k) {
//			for i := 0; i < 7-int(k); i++ {
//				colorArr = append(colorArr, int(k))
//			}
//			subNum -= 7 - int(k)
//		} else {
//			for i := 0; i < subNum; i++ {
//				colorArr = append(colorArr, int(k))
//				subNum = 0
//			}
//		}
//	}
//	//将二次中奖最终花色填充
//	idx := 0
//	for k, v := range newGrids {
//		for k1, v1 := range v {
//			if v1 == 0 && idx < len(colorArr) {
//				newGrids[k][k1] = int32(colorArr[idx])
//				idx++
//			}
//		}
//	}
//	return newGrids
//}
//
//// GetColor 取花色
//func GetColor(color int) int {
//	colorMap := make(map[int]int)
//	colorMap[1] = 1001
//	colorMap[2] = 1002
//	colorMap[3] = 1003
//	colorMap[4] = 1004
//	colorMap[5] = 1005
//	colorMap[6] = 1006
//	colorMap[7] = 1007
//	colorMap[8] = 1008
//	colorMap[9] = 1009
//	colorMap[10] = 1010
//	colorMap[11] = 1011
//	colorMap[12] = 1012
//	colorMap[13] = 1013
//	colorMap[14] = 1014
//	colorMap[15] = 1015
//	if color > 0 {
//		delete(colorMap, color)
//	}
//	for _, v := range colorMap {
//		return v
//	}
//	return 0
//}
//
//// FillInGrids 将最终数据填充到二维数组
//func FillInGrids(grids map[int]int, oldGrids [5][6]int32) [5][6]int32 {
//	idx := 0
//	for i := 0; i < 5; i++ {
//		for j := 0; j < 6; j++ {
//			if oldGrids[i][j] == 0 {
//				oldGrids[i][j] = int32(grids[idx])
//				idx++
//			}
//		}
//	}
//	return oldGrids
//}
//
//// DelColor 删除指定花色重新排列
//func DelColor(grids [5][6]int32) [5][6]int32 {
//	//重新排列
//	for i := 4; i > -1; i-- {
//		for j := 0; j < 6; j++ {
//			if grids[i][j] == 0 && i > 0 && grids[i-1][j] > 0 {
//				grids[i][j] = grids[i-1][j]
//				grids[i-1][j] = 0
//			}
//		}
//	}
//	return grids
//}
//
//func DoubleArrToMap(girds [5][6]int32) map[int32]int32 {
//	firMap := make(map[int32]int32)
//	i := 1
//	for _, v := range girds {
//		for _, v1 := range v {
//			firMap[int32(i)] = v1
//			i++
//		}
//	}
//	return firMap
//}
//
//func MapToDoubleArr(data map[int32]int32) [5][6]int32 {
//	arr := [5][6]int32{}
//	i := 0
//	for _, v := range data {
//		arr[i%5][i%6] = v
//		i++
//	}
//	return arr
//}
