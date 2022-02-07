package model

type GoodsSpec struct {
	Id       int    `xorm:"not null pk INT(11)"`
	SpecName string `xorm:"not null comment('规格名称例如（内存、颜色）') VARCHAR(255)"`
	SpecNo   string `xorm:"not null comment('规格编号') VARCHAR(255)"`
}
