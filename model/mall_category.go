package models

type MallCategory struct {
	CategoryId   int    `xorm:"not null pk INT(11)"`
	CategoryName string `xorm:"VARCHAR(255)"`
}
