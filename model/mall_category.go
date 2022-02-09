package model

// MallCategory
// 平台商品类目
type MallCategory struct {
	ID   int    `xorm:"'category_id'"`
	Name string `xorm:"'category_name'"`
}
