package db

import (
	"gorm.io/gorm"
)

type Chapter struct {
	gorm.Model
	Pos   string `gorm:"column:pos;type:text"`
	Color string `gorm:"column:color;type:text"`
	Types string `gorm:"column:types;type:text"`
}

func (c *Chapter) TableName() string {
	return "ball"
}
