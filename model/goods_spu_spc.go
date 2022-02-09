package model

type GoodsSpuSpc struct {
	Id int `xorm:"not null pk comment('主键') INT(11)"`
	SpuId  int `xorm:"not null comment('SPU的id') INT(11)"`
	SpecId int `xorm:"not null comment('规格的id') INT(11)"`
}
