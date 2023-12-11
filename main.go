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
	//http.InitHttpService("192.168.204.67:9999")
}
