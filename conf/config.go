package conf

import (
	"github.com/spf13/viper"
)

type Log struct {
	FileName string
	LogLevel string
	FilePath string
}

type Mysql struct {
	Ip       string
	Username string
	Password string
	Database string
}
type Config struct {
	Log   Log
	Mysql Mysql
}

var Conf Config

func ConfigInit() {
	v := viper.New()
	v.AddConfigPath("./")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.ReadInConfig()
	v.Unmarshal(&Conf)
}
