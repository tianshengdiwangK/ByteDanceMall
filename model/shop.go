package model

import (
	"time"
)

type Shop struct {
	ShopId          int       `xorm:"not null pk autoincr comment('shop id') INT(11)"`
	SellerId        int       `xorm:"not null comment('店铺所属卖家id') INT(11)"`
	ShopDescription string    `xorm:"not null comment('店铺简介') VARCHAR(255)"`
	ShopName        string    `xorm:"not null comment('店铺名称') VARCHAR(255)"`
	ShopCreateTime  time.Time `xorm:"not null comment('店铺建立日期') DATETIME"`
}
