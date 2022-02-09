package model

import (
	"time"
)

type Seller struct {
	SellerId    int    `xorm:"not null pk comment('商家id') INT(11)"`
	SellerEmail string `xorm:"not null comment('商家email') VARCHAR(255)"`
	SellerPhone    string    `xorm:"not null comment('商家电话') VARCHAR(255)"`
	SellerUsername string    `xorm:"not null comment('商家用户名') VARCHAR(255)"`
	SellerPassword string    `xorm:"not null comment('商家密码') VARCHAR(255)"`
	CreateTime     time.Time `xorm:"not null comment('创建时间') DATETIME(6)"`
}
