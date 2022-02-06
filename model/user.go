package model
// Xorm Struct
type Users struct {
	Username string `xorm:"VARCHAR(255)"`
	Password string `xorm:"VARCHAR(255)"`
}

