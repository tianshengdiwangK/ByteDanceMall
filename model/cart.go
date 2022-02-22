package model

type Cart struct {
	CartId int    `xorm:"not null pk comment('购物车 id') INT(11)"`
	UserId int    `xorm:"not null comment('购物车所属用户id') index INT(11)"`
	SkuId  int    `xorm:"not null comment('产品id') INT(11)"`
	Count  int    `xorm:"not null comment('商品数量') INT(11)"`
	Price  string `xorm:"not null comment('产品价格') DECIMAL(10,2)"`
	Name   string `xorm:"not null comment('产品名称') VARCHAR(255)"`
}
