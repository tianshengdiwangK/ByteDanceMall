package model

type GoodsCategory struct {
	CategoryId   int    `xorm:"not null pk comment('种类id') INT(11)"`
	CategoryName string `xorm:"not null comment('种类名称') VARCHAR(255)"`
	CategoryIcon        string `xorm:"not null comment('种类图标') VARCHAR(255)"`
	CategoryDescription string `xorm:"not null comment('种类描述') TEXT"`
	CategoryShowStatus  int    `xorm:"not null comment('是否展示该种类') TINYINT(1)"`
	ProductCount int `xorm:"not null comment('该种类下产品数量') INT(11)"`
	ShopId       int `xorm:"not null comment('店铺id') index INT(11)"`
	ParentId     int `xorm:"not null comment('父分类id值，0表示最高级分类') INT(11)"`
}
