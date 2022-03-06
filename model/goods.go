package model

type Goods struct {
	Id              int    `xorm:"not null pk autoincr comment('商品id') INT"`
	BrandId         int    `xorm:"not null comment('品牌id') INT" json:"brand_id" binding:"required"`
	GoodsCategoryId int    `xorm:"not null comment('店铺类目id') INT" json:"goods_category_id" binding:"required"`
	Name            string `xorm:"not null comment('名称') VARCHAR(45)" binding:"required"`
	GoodsSn         string `xorm:"comment('货号') VARCHAR(45)" json:"goods_sn"`
	Price           string `xorm:"comment('售价') VARCHAR(45)"`
	Description     string `xorm:"comment('介绍') VARCHAR(45)"`
	Stock           int    `xorm:"comment('库存') INT"`
	ShopId       	int `xorm:"not null comment('店铺id') index INT(11)"`
}
