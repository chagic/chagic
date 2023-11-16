package conf

import (
	"chagic/model"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	MySQL  MySQLConfig
	Log    LogConfig
	Server ServerConfig
}

type MySQLConfig struct {
	Host        string
	Name        string
	Password    string
	Port        string
	TablePrefix string
	User        string
}

type LogConfig struct {
	Path  string
	Level string
}

type ServerConfig struct {
	Mode string
	Port string
}

var c Config

func init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	viper.Unmarshal(&c)
	DbUser := GetConfig().MySQL.User
	DbPassWord := GetConfig().MySQL.Password
	DbHost := GetConfig().MySQL.Host
	DbPort := GetConfig().MySQL.Port
	DbName := GetConfig().MySQL.Name
	TablePrefix := GetConfig().MySQL.TablePrefix

	dsn := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(dsn, TablePrefix)
}

func GetConfig() Config {
	return c
}
