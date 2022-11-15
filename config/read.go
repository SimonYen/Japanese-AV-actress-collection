package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func set() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(nil)
	}
}

func (receiver *ServerConfig) Read() {
	set()
	receiver.Address = viper.GetString("server.address")
	receiver.Port = viper.GetInt("server.port")
}

func (receiver *ServerConfig) ToString() string {
	return fmt.Sprintf("%s:%d", receiver.Address, receiver.Port)
}

func (receiver *MysqlConfig) Read() {
	set()
	receiver.Address = viper.GetString("mysql.address")
	receiver.Database = viper.GetString("mysql.database")
	receiver.Password = viper.GetString("mysql.password")
	receiver.Port = viper.GetInt("mysql.port")
	receiver.User = viper.GetString("mysql.user")
}
func (receiver *MysqlConfig) ToString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", receiver.User, receiver.Password, receiver.Address, receiver.Port, receiver.Database)
}
