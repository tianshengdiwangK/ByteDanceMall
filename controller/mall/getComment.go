package mall

import (
	"fmt"
	"log"
	"login_register_demo/config"
	"login_register_demo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodsComment struct {
	Id        int    `xorm:"not null pk comment('评论id') INT(11)"`
	ProductId int    `xorm:"not null pk comment('评论所属产品id') INT(11)"`
	UserId    int    `xorm:"not null comment('评论人') INT(11)"`
	Content   string `xorm:"not null comment('评论内容') VARCHAR(255)"`
}

func GetGoodsComment(c *gin.Context) {
	var GoodsComments []GoodsComment
	productId := c.Request.URL.Query().Get("product_id")
	fmt.Println(productId)
	err := config.Engine.Table("goods_comment").Where("product_id=?", productId).Find(&GoodsComments)
	if err != nil {
		log.Println("get goodsComment:", err)
	}
	res := make([][]string, 0)
	for _, v := range GoodsComments {
		st2 := new(model.User)
		config.Engine.Table("user").Where("user_id=?", v.UserId).Get(st2)
		cur := []string{"commentcontent:" + v.Content, "userid:" + strconv.Itoa(v.UserId), "username:" + st2.UserUsername}
		res = append(res, cur)
	}
	c.JSON(http.StatusOK, res)
}
