package database

import (
	"jyoyuu/config"
	"jyoyuu/models"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var DB *xorm.Engine

func Connect() {
	mc := new(config.MysqlConfig)
	mc.Read()
	eng, err := xorm.NewEngine("mysql", mc.ToString())
	if err != nil {
		panic(err)
	}
	DB = eng
}

func Migrate() {
	err := DB.Sync2(new(models.User))
	if err != nil {
		panic(err)
	}
}
