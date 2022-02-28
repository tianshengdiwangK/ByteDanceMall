package model

type GoodsImage struct {
	GoodsId int    `xorm:"not null comment('商品id') index INT"`
	Id      int    `xorm:"not null pk autoincr comment('商品图片id') INT"`
	Image   string `xorm:"not null comment('图片相对路径') VARCHAR(255)"`
}
