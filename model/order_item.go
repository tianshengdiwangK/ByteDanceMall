package model

import (
	"time"
)

type OrderItem struct {
	OrderItemId      int       `xorm:"not null pk comment('id值') INT(11)"`
	OrderId          int       `xorm:"not null comment('订单id') INT(11)"`
	SkuId            int       `xorm:"not null comment('商品id') INT(11)"`
	ShopId           int       `xorm:"not null comment('订单所属店铺id') INT(11)"`
	CurrentUnitPrice string    `xorm:"not null comment('生成订单时的商品单价，单位是元，保留两位小数') DECIMAL(10,2)"`
	ProductCount     int       `xorm:"not null comment('产品数量') INT(11)"`
	CreateTime       time.Time `xorm:"not null comment('订单创建时间') DATETIME"`
	Consignee        string    `xorm:"not null comment('收货人') VARCHAR(255)"`
	Address          string    `xorm:"not null comment('收货地址') VARCHAR(255)"`
	Phone            string    `xorm:"not null comment('收货电话') VARCHAR(255)"`
}
