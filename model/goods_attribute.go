package model

type GoodsAttribute struct {
	Id                       int    `xorm:"not null INT(11)"`
	GoodsAttributeCategoryId int    `xorm:"INT(11)"`
	Name                     string `xorm:"VARCHAR(45)"`
	SelectType               int 	`xorm:"INT(11)"`
	InputList                string `xorm:"VARCHAR(45)"`
}
