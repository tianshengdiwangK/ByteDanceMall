package cart

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"login_register_demo/model"
	"login_register_demo/utils/errmsg"
	"net/http"
	"strconv"
)

type ProductInfo struct {
	UserId      int    `json:"user_id"`
	ProductId   int    `json:"product_id"`
	ShopId      string `json:"shop_id"`      //字段未用上
	ProductType string `json:"product_type"` //字段未用上
	ProductNum  int    `json:"product_num"`
}

var code int
var err error

//添加购物车商品
func CartAddProduct(c *gin.Context) {
	var ProInfo ProductInfo
	err = c.ShouldBindJSON(&ProInfo)
	if err != nil {
		code = errmsg.ERROR_JSON_TYPE_WRONG
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		goodsSku := model.GoodsSku{}
		has, err := config.Engine.Id(ProInfo.ProductId).Get(&goodsSku)
		if err != nil && has == false {
			code = errmsg.ERROR_DATABASE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Message": errmsg.GetErrMsg(code),
			})
		}
		cart := model.Cart{}
		has, err = config.Engine.Where("user_id=?", ProInfo.UserId).And("sku_id=?", ProInfo.ProductId).Get(&cart)
		price, _ := strconv.ParseFloat(goodsSku.Price, 64)
		if has == false {
			goodsSpu := model.GoodsSpu{}

			config.Engine.Where("id=?", goodsSku.SpuId).Get(&goodsSpu)
			cart = model.Cart{
				UserId: ProInfo.UserId,
				SkuId:  ProInfo.ProductId,
				Count:  ProInfo.ProductNum,
				Price:  strconv.FormatFloat(float64(ProInfo.ProductNum)*price, 'f', 2, 64),
				Name:   goodsSpu.ProductName,
			}
			_, err = config.Engine.Insert(&cart)
			if err != nil {
				code = errmsg.ERROR_DATABASE_WRONG
				c.JSON(http.StatusOK, gin.H{
					"Status":  code,
					"Message": errmsg.GetErrMsg(code),
				})
			}
			code = errmsg.SUCCSE
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Message": errmsg.GetErrMsg(code),
			})
		} else {
			cart.Count += ProInfo.ProductNum
			cart.Price = strconv.FormatFloat(float64(cart.Count)*price, 'f', 2, 64)
			config.Engine.Id(cart.CartId).Cols("count").Cols("price").Update(&cart)
			code = errmsg.SUCCSE
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Message": errmsg.GetErrMsg(code),
			})
		}
	}

}
