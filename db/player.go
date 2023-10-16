package db

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Uid        string `gorm:"column:uid;type:varchar(30);default:err;NOT NULL;unique"`
	Name       string `gorm:"column:name;type:varchar(20)"`
	ChangeFree int    `gorm:"column:change_free;type:int(11);default:1"`
	Avatar     string `gorm:"column:avatar;type:varchar(30)"`
	Level      int32  `gorm:"column:level;type:int(11);default:1;NOT NULL"`
	LevelNum   int32  `gorm:"column:level_num;type:int(11);default:0;NOT NULL"`
	Gold       int32  `gorm:"column:gold;type:int(11);default:0;NOT NULL"`
	Diamond    int32  `gorm:"column:diamond;type:int(11);default:0;NOT NULL"`
	LoginDays  string `gorm:"column:login_days;type:text"`
	RedeemCode string `gorm:"column:redeem_code;type:text"`
}

func (player *Player) TableName() string {
	return "player"
}
