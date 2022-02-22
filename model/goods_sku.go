package models

type GoodsSku struct {
	Id int `xorm:"not null pk comment('主键
') INT(11)"`
	GoodsId int    `xorm:"comment('SPU的id') INT(11)"`
	SkuCode string `xorm:"comment('店铺id') VARCHAR(84)"`
	Price   string `xorm:"comment('售价') DECIMAL(10,2)"`
	Stock   string `xorm:"comment('sku编号') VARCHAR(255)"`
	Pic     string `xorm:"VARCHAR(255)"`
	SpData  string `xorm:"VARCHAR(255)"`
}
