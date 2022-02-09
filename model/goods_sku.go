package model

type GoodsSku struct {
	Id int `xorm:"not null pk comment('主键') INT(11)"`
	SpuId  int    `xorm:"not null comment('SPU的id') INT(11)"`
	ShopId int    `xorm:"not null comment('店铺id') INT(11)"`
	Price  string `xorm:"not null comment('售价') DECIMAL(10,2)"`
	SkuNo  string `xorm:"not null comment('sku编号') VARCHAR(255)"`
	Stock  int    `xorm:"not null comment('商品库存') INT(11)"`
}
