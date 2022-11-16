package models

import "time"

type User struct {
	Id       int64
	Name     string `xorm:"notnull"`
	Phone    string `xorm:"unique"`
	Password string
	Created  time.Time `xorm:"created"`
}
