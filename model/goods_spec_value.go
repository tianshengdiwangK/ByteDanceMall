package model

type GoodsSpecValue struct {
	Id        int    `xorm:"not null pk INT(11)"`
	SpecId    int    `xorm:"not null comment('规格id值') INT(11)"`
	SpecValue string `xorm:"not null comment('具体的规格内容如（32G、黑色）') VARCHAR(255)"`
}
