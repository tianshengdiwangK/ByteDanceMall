package models

type OrderItem struct {
	Id               int    `xorm:"not null pk comment('id值') INT(11)"`
	ShopId           int    `xorm:"comment('订单所属店铺id') INT(11)"`
	CurrentUnitPrice string `xorm:"comment('生成订单时的商品单价，单位是元，保留两位小数') DECIMAL(10,2)"`
	ProductCount     int    `xorm:"comment('产品数量') INT(11)"`
	TotalPrice       string `xorm:"DECIMAL(10,2)"`
	Brand            string `xorm:"VARCHAR(45)"`
	Name             string `xorm:"VARCHAR(45)"`
	SpData           string `xorm:"VARCHAR(45)"`
	GoodsSn          string `xorm:"VARCHAR(45)"`
	OrderId          int    `xorm:"INT(11)"`
}
