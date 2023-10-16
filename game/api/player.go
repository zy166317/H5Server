package api

import "leaf_server/db"

type Player struct {
	DbPlayer   *db.Player
	LoginDays  []int32         //连续登录
	RedeemCode map[string]bool //兑换码
	BackPack   map[int32]int32 //道具背包
}

func (p *Player) CreateRole() {

}