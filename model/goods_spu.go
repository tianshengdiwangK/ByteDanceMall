package model

type GoodsSpu struct {
	Id         int    `xorm:"not null pk INT(11)"`
	BrandId    int    `xorm:"not null comment('品牌id') INT(11)"`
	CategoryId int    `xorm:"not null comment('种类id') INT(11)"`
	SpuNo      string `xorm:"not null comment('商品编号，唯一') VARCHAR(255)"`
	ProductName string `xorm:"not null comment('商品名称') VARCHAR(255)"`
	ProductDescription string `xorm:"comment('商品介绍') VARCHAR(255)"`
}
