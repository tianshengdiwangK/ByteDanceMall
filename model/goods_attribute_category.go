package model

type GoodsAttributeCategory struct {
	Id                  int    `xorm:"not null INT(11)"`
	GoodsAttributeName  string `xorm:"VARCHAR(45)" json:"goods_attribute_name"`
	GoodsAttributeCount int    `xorm:"INT(11)"`
	ShopId       		int `xorm:"not null comment('店铺id') index INT(11)" json:"shop_id"`
}
