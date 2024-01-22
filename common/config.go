package common

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/recordfile"
	"leaf_server/gamedata"
	"sync"
)

type Config struct {
	cf   *recordfile.RecordFile // 配置文件
	Lock sync.RWMutex           // cf锁
}

var (
	configs    map[int]*Config     // 配置map
	configType map[int]interface{} // 配置注册map
)

// 表名配置id
const (
	CHAPTER_STARS = iota
	ITEM
	LOOT
	BALL
	BAFFLE
	LOGIN_SEVEN_DAY
)

// 配置注册初始化
// TODO
func ConfigInit() {
	configs = make(map[int]*Config)
	configType = make(map[int]interface{})
	RegisterCf(CHAPTER_STARS, gamedata.Chapter_stars{})
	RegisterCf(ITEM, gamedata.Item{})
	RegisterCf(LOOT, gamedata.Loot{})
	RegisterCf(BALL, gamedata.Ball{})
	RegisterCf(BAFFLE, gamedata.Baffle{})
	RegisterCf(LOGIN_SEVEN_DAY, gamedata.Login_seven_day{})
}

// @desc: 获取配置文件
// @args:[{cfId: 配置id}]
// @return: [{cf: 配置文件}]
// @alarm:
func GetCf(cfId int) (cf *recordfile.RecordFile) {
	if v, ok := configs[cfId]; ok {
		v.Lock.RLock()
		defer v.Lock.RUnlock()
		return v.cf
	}
	return nil
}

// @desc: 设置配置文件
// @args:[{cfId: 配置id}, {cf: 配置文件}]
// @return: []
// @alarm:
func SetCf(cfId int, cf *recordfile.RecordFile) {
	if v, ok := configs[cfId]; ok {
		v.Lock.Lock()
		defer v.Lock.Unlock()
		v.cf = cf
	} else {
		configs[cfId] = &Config{
			cf: cf,
		}
	}
}

// @desc: 重载特定的配置文件
// @args:[{cfId: 配置id}]
// @return: []
// @alarm:
func Reload(cfId int) {
	// 错误hold
	defer func() {
		if v := recover(); v != nil {
			log.Error("cfId relaod fail: %v", cfId)
			return
		}
	}()
	if v, ok := configType[cfId]; ok {
		Rf := readRf(v)
		SetCf(cfId, Rf)
	}
}

// @desc: 重载所有配置文件
// @args:[]
// @return: []
// @alarm:
func ReloadAll() {
	for cfId, _ := range configs {
		Reload(cfId)
		log.Release("Reload config file id: %d", cfId)
	}
}

// @desc: 注册相应配置文件
// @args:[{cfId: 配置id}, {st: 配置结构体}]
// @return: []
// @alarm:
func RegisterCf(cfId int, st interface{}) {
	configType[cfId] = st
	Rf := readRf(st)
	SetCf(cfId, Rf)
}
