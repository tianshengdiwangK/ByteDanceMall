package model

type GoodsSkuSpecValue struct {
	Id          int `xorm:"not null pk comment('主键') INT(11)"`
	SpecValueId int `xorm:"not null comment('规格值id') INT(11)"`
	SkuId       int `xorm:"not null comment('sku_id') INT(11)"`
}
