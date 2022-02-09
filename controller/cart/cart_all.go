package cart

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"login_register_demo/model"
	"net/http"
	"strconv"
)

// GetCartAll GET /cart/all
// 获取用户购物车全部商品
func GetCartAll(c *gin.Context) {
	userid, ok := c.GetQuery("userid")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg " : "参数错误",
		})
	}

	type queryRes struct {
		model.Cart	`xorm:"extends"`
		model.GoodsSku	`xorm:"extends"`
		model.GoodsSpu	`xorm:"extends"`
		model.Shop	`xorm:"extends"`
		model.GoodsCategory	`xorm:"extends"`
	}

	r := make([]queryRes, 0)

	config.Engine.Table("cart").
		Join("INNER", "goods_sku", "goods_sku.sku_no = cart.sku_id").
		Join("INNER", "goods_spu", "goods_spu.spu_no = goods_sku.spu_id").
		Join("INNER", "shop", "goods_sku.shop_id = shop.shop_id").
		Join("INNER", "goods_category", "goods_category.category_id = goods_spu.category_id").
		Where("cart.user_id = ?", userid).
		Find(&r)

	jsonRes := make(map[string][][]string)

	for _, val := range r {
		t := []string{val.Cart.Name, val.Cart.Price, val.GoodsCategory.CategoryName, strconv.Itoa(val.Cart.Count)}
		jsonRes[val.Shop.ShopName] = append(jsonRes[val.Shop.ShopName], t)
	}


	c.JSON(http.StatusOK, jsonRes)
}