package db

import (
	"fmt"
	"time"

	"gorm.io/gorm/schema"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getConnStr(user string, pass string, protocol string, addr string, db string) string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, pass, protocol, addr, db)
}

// func InitDB(connStr string, opts ...Option) {
func ConnectToMysql(mycfg *Mysql) (*gorm.DB, error) {
	connStr := getConnStr(mycfg.User, mycfg.Password, mycfg.Protocol, mycfg.Address, mycfg.DBName)
	dbc, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	myDB, err := dbc.DB()
	if err != nil {
		return nil, err
	}
	err = myDB.Ping()
	if err != nil {
		return nil, err
	}

	dbc.Debug()
	//dbc.AutoMigrate()
	myDB.SetConnMaxLifetime(time.Duration(mycfg.ConnMaxLifeTime) * time.Second)
	myDB.SetMaxOpenConns(mycfg.MaxOpenConnNum)
	myDB.SetMaxIdleConns(mycfg.MaxIdleConnNum)
	return dbc, nil
}
