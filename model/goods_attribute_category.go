package model

type GoodsAttributeCategory struct {
	Id                  int    `xorm:"not null INT(11)"`
	GoodsAttributeName  string `xorm:"VARCHAR(45)"`
	GoodsAttributeCount int    `xorm:"INT(11)"`
	ShopId       		int `xorm:"not null comment('店铺id') index INT(11)"`
}
