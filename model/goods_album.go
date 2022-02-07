package model

type GoodsAlbum struct {
	Id int `xorm:"not null pk comment('主键') INT(11)"`
	SpuId  int    `xorm:"not null comment('图片所属商品id') INT(11)"`
	PicDes string `xorm:"not null comment('图片描述') VARCHAR(255)"`
	PicUrl string `xorm:"comment('图片所属url') VARCHAR(255)"`
}
