package model

import (
	"time"
)

type GoodsComment struct {
	CommentId int `xorm:"not null pk comment('评论id') INT(11)"`
	SpuId     int `xorm:"not null pk comment('评论所属产品id') INT(11)"`
	UserId    int `xorm:"not null comment('评论人') INT(11)"`
	CreateTime time.Time `xorm:"not null comment('评论创建时间') DATETIME"`
	Pics       string    `xorm:"not null comment('评论图片信息') VARCHAR(255)"`
	Content    string    `xorm:"not null comment('评论内容') VARCHAR(255)"`
}
