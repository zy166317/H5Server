package db

import (
	"github.com/name5566/leaf/log"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Uid              string `gorm:"column:uid;type:varchar(30);default:err;NOT NULL;unique"`
	BackPack         string `gorm:"column:back_pack;type:text"` //球球皮肤背包
	Gold             int32  `gorm:"column:gold;type:int(11);default:0;NOT NULL"`
	CheckPointRecord string `gorm:"column:check_point_record;type:text"` //通过关卡信息记录
	LoginDays        string `gorm:"column:login_days;type:text"`
}

func (p *Player) TableName() string {
	return "player"
}

func GetPlayerInfo(uid string) (*Player, error) {
	player := &Player{}
	err := DBC.Where("uid = ? ", uid).First(player).Error
	if err != nil {
		log.Release("get player info error: %v", err.Error())
		return nil, err
	}
	return player, nil
}
