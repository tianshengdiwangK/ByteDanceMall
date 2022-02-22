package models

type GoodsAttributeCategory struct {
	Id                  int    `xorm:"not null INT(11)"`
	GoodsAttributeName  string `xorm:"VARCHAR(45)"`
	GoodsAttributeCount int    `xorm:"INT(11)"`
}
