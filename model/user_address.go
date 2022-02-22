package model

type UserAddress struct {
	Id       int    `xorm:"not null pk comment('主键') INT(11)"`
	UserId   int    `xorm:"not null comment('用户id') INT(11)"`
	Province string `xorm:"not null comment('省') VARCHAR(255)"`
	City     string `xorm:"not null comment('市') VARCHAR(255)"`
	Area     string `xorm:"not null comment('区') VARCHAR(255)"`
	People   string `xorm:"not null comment('联系人') VARCHAR(255)"`
	Phone    string `xorm:"not null comment('手机号') VARCHAR(255)"`
}
