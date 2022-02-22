package model

type Goods struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	BrandId         string `xorm:"VARCHAR(45)"`
	GoodsCategoryId string `xorm:"VARCHAR(45)"`
	Name            string `xorm:"VARCHAR(45)"`
	Pic             string `xorm:"VARCHAR(255)"`
	GoodsSn         string `xorm:"VARCHAR(45)"`
	Price           string `xorm:"VARCHAR(45)"`
	Desription      string `xorm:"VARCHAR(45)"`
	Stock           string `xorm:"VARCHAR(45)"`
}
