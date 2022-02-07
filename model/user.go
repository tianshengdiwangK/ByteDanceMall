package model

import (
	"time"
)

type User struct {
	UserId    int    `xorm:"not null pk autoincr comment('用户id') INT(11)"`
	UserEmail string `xorm:"not null comment('用户email') VARCHAR(255)"`
	UserPhone string `xorm:"not null comment('用户注册时的手机号') VARCHAR(255)"`
	UserUsername string `xorm:"not null comment('用户名') VARCHAR(255)"`
	UserPassword string `xorm:"not null comment('用户密码') VARCHAR(255)"`
	Status       int    `xorm:"not null comment('记录用户是否登录') INT(11)"`
	CreateTime time.Time `xorm:"not null comment('记录用户创建时间') DATETIME"`
	Icon string `xorm:"not null comment('用户头像') VARCHAR(255)"`
}

