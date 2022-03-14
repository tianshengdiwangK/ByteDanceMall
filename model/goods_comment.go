package model

type GoodsComment struct {
	CommentId int `xorm:"not null pk comment('评论id') INT(11)"`
	ProductId     int `xorm:"not null pk comment('评论所属产品id') INT(11)"`
	UserId    int `xorm:"not null comment('评论人') INT(11)"`
	Content    string    `xorm:"not null comment('评论内容') VARCHAR(255)"`
}
