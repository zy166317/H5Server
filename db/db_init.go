package db

import "gorm.io/gorm"

var DBC *gorm.DB

func Init(mycfg *Mysql) error {
	var err error
	DBC, err = ConnectToMysql(mycfg)
	if err != nil {
		return err
	}

	return autoMigrate()
}

func autoMigrate() error {
	return DBC.AutoMigrate(&Player{})
}
