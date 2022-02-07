package model

import (
	"time"
)

type Order struct {
	OrderId int `xorm:"not null pk comment('订单id') INT(11)"`
	OrderNo int `xorm:"not null comment('如果对编号格式没什么要求，可使用雪花算法https://blog.csdn.net/u011499747/article/details/78254990来生成') INT(11)"`
	UserId  int `xorm:"not null comment('订单所属客户id') INT(11)"`
	ShopId          int       `xorm:"not null comment('订单所属店铺id') INT(11)"`
	OrderStatus     int       `xorm:"not null comment('订单的状态，1表示已完成，0表示未完成。') INT(11)"`
	OrderCreateTime time.Time `xorm:"not null comment('订单创建信息') DATETIME"`
	OrderPrice      string    `xorm:"not null comment('订单金额') DECIMAL(10,2)"`
}
