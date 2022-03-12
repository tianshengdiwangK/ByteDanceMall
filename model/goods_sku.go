package model

import "strings"

type GoodsSku struct {
	Id      int    `xorm:"not null pk comment('主键') INT(11)"`
	GoodsId int    `xorm:"not null comment('SPU的id') INT(11)" json:"goods_id"`
	SkuCode string `xorm:"comment('sku编号') VARCHAR(84)" json:"sku_code"`
	Price   string `xorm:"not null comment('售价') DECIMAL(10,2)" binding:"required"`
	Stock   int    `xorm:"comment('库存') INT" binding:"required,gte=0"`
	SpData  []Attr `xorm:"comment('销售属性') JSON" json:"sp_data"`
}

// Attr 销售属性，以JSON格式存在goods_sku表中
type Attr struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GetBriefAttr 返回销售属性的简短字符串
//  eg. "白色 | 256GB"
func (sku *GoodsSku) GetBriefAttr() string {
	attrs := make([]string, len(sku.SpData))
	for index, attr := range sku.SpData {
		attrs[index] = attr.Value
	}
	return strings.Join(attrs, " | ")
}
