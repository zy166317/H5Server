package api

import (
	"github.com/name5566/leaf/log"
	"github.com/robfig/cron/v3"
)

var (
	CronTimer *cron.Cron
	CronId    = make(map[string][]cron.EntryID)
)

//定时任务Cron包
//玩家开启的定时任务时同步将任务id存入全局
//玩家下线同步移除玩家定时任务

func CronInit() {
	c := cron.New(cron.WithSeconds())
	CronTimer = c
	//添加定时任务
	c.AddFunc("", func() {

	})
	c.Run()
}

// AddPlayerCronTask 添加玩家个人定时任务
func AddPlayerCronTask(uid, spec string, cmd func()) {
	entryID, err := CronTimer.AddFunc(spec, cmd)
	if err != nil {
		log.Error("添加定时任务失败")
		return
	}
	CronId[uid] = append(CronId[uid], entryID)
}

// DelPlayerCronTask 删除玩家个人定时任务
func DelPlayerCronTask(uid string) {
	if _, has := CronId[uid]; has {
		delete(CronId, uid)
	}
}
