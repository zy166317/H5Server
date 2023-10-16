package common

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/recordfile"
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
	GLOBAL_CONFIG = iota
	ATTRIBUTE
	BLUE_PRINT
	CHALLENGE_BOSS
	CHAPTER
	COMBAT_MODEL
	EFFECT_MODEL
	EQUIP_RANDOM
	FLIGHT_MAP
	GEM
	GEM_RULE
	ITEM
	KUNGFU
	MAIN_TASK
	MONSTER
	MYTHICS
	MYTHICS_GROUP
	MYTHICS_MAP
	NATURE_TALISMAN
	NATURE_TALISMAN_POWER
	NATURE_TALISMAN_UPGRADE
	REALM
	ROLE_LEVEL
	SKILL_MODEL
	SOUND_EFFECT
	STORY_MAIN
	STORY_RANDOM
	STORY_RHYTHM
	TALLY
	TALLY_ATLAS
	TALLY_BAIT
	TALLY_BOX
	TALLY_EQUIP
	TALLY_FETTER
	TALLY_ORDER_TASK
	TALLY_TASK
	TALLY_TASK_GROUP
	TALLY_TYPE
	TRIPOD_BOX
	CORE_TEST
	REALM_TASK
	GRASS
	PRESCRIPT
	SHOP_GOODS
	ALCHEMY
	TALLY_SCENE
	ARENA_MATE
	ARENA_INTEGRAL
	ARENA_INTEGRAL2
	AREBA_DAILY_RANK_REWARD
	ARENA_RANK_REWARD
	SHOP
	ACTIVITY_TASKS
	LOOT
	DIARY
	MARKET
)

// 配置注册初始化
// TODO
func ConfigInit() {
	configs = make(map[int]*Config)
	configType = make(map[int]interface{})
	//RegisterCf(GEM, gamedata.Gem{})
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
