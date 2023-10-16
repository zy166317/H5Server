package db

import "gorm.io/gorm/logger"

type Mysql struct {
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	Protocol        string `yaml:"protocol"`
	Address         string `yaml:"address"`
	DBName          string `yaml:"dbName"`
	Driver          string `yaml:"driver"`
	ConnMaxLifeTime int    `yaml:"connMaxLifeTime"`
	MaxIdleConnNum  int    `yaml:"maxIdleConnNum"`
	MaxOpenConnNum  int    `yaml:"maxOpenConnNum"`
}

type Redis struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Password     string `yaml:"password"`
	DBIndex      int    `yaml:"db"`
	MinIdleConns int    `yaml:"minIdleConns"` //最小空闲连接数
	IdleTimeout  int    `yaml:"idleTimeout"`  //多久没有使用回收
	MaxConnAge   int    `yaml:"maxConnAge"`   //连接生命周期
	PoolSize     int    `yaml:"poolSize"`
}

type Configuration struct {
	ListenAddr string        `yaml:"listenAddr"`
	ServerID   uint16        `yaml:"serverID"`
	DBNo       int           `yaml:"dbNo"`
	Log        logger.Config `yaml:"log"`
	DBGame     Mysql         `yaml:"dbGame"`
	DBRedis    Redis         `yaml:"dbRedis"`
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}
