package db

import "time"

type Chapter struct {
	Id        int64     `gorm:"column:id;type:bigint(20);primaryKey;unique;AUTO_INCREMENT"`
	Pos       string    `gorm:"column:pos;type:text"`
	Color     string    `gorm:"column:color;type:text"`
	Types     string    `gorm:"column:types;type:text"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3)"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3)"`
}

func (c *Chapter) TableName() string {
	return "ball"
}
