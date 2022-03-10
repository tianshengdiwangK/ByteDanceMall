package model

type Cart struct {
	Id        int    `xorm:"not null pk autoincr comment('购物车项 id') INT"`
	UserId    int    `xorm:"not null comment('购物车所属用户id') index INT"`
	Name      string `xorm:"not null comment('产品名称') VARCHAR(255)"`
	Count     int    `xorm:"not null comment('商品数量') INT"`
	ProductId int    `xorm:"not null comment('商品id') INT"`
	SkuId     int    `xorm:"not null comment('sku id') INT"`
	ShopName  string `xorm:"not null comment('店铺名') VARCHAR(255)"`
}
