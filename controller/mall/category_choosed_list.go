package mall

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"login_register_demo/model"
	"net/http"
)

func GetCategoryGoods(c *gin.Context) {
	categoryId, ok := c.GetQuery("category_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"msg ":    "请求参数错误",
		})
	}

	type goodsItem struct {
		model.Goods      `xorm:"extends"`
		model.GoodsImage `xorm:"extends"`
	}
	res := make([]goodsItem, 0)
	config.Engine.Table("goods").
		Join("INNER", "goods_image", "goods_image.goods_id = cart.product_id and goods_image.is_primary = 1").
		Where("goods.goods_category_id = ?", categoryId).
		Find(&res)

	//封装结果
	type queryRes struct {
		productid   int    `json:"productid"`
		productname string `json:"productname"`
		productinfo string `json:"productinfo"`
		price       string `json:"price"`
		productimg  string `json:"productimg"`
	}
	r := make([]queryRes, len(res))
	for i, val := range res {
		r[i].productid = val.GoodsId
		r[i].productname = val.Name
		r[i].productinfo = val.Description
		r[i].price = val.Price
		r[i].productimg = val.Image
	}
	c.JSON(http.StatusOK, r)
}
