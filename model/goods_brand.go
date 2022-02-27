package model

type GoodsBrand struct {
	Id int `xorm:"not null pk comment('主键') INT(11)"`
	BrandName string `xorm:"not null comment('品牌名称') VARCHAR(255)"`
}
