package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"login_register_demo/config"
	"login_register_demo/model"
)

// GetCartAll
//  GET /cart/all
// 获取用户购物车全部商品
func GetCartAll(c *gin.Context) {
	// 从context中获取user_id
	val, _ := c.Get("user_id")
	userId := val.(int) // 断言为int

	type CartItem struct {
		model.Cart       `xorm:"extends"`
		model.GoodsSku   `xorm:"extends"`
		model.GoodsImage `xorm:"extends"`
	}

	cartItems := make([]CartItem, 0)

	config.Engine.Table("cart").
		Join("INNER", "goods_sku", "goods_sku.id = cart.sku_id").
		Join("INNER", "goods_image", "goods_image.goods_id = cart.product_id and goods_image.is_primary = 1").
		Where("cart.user_id = ?", userId).
		Find(&cartItems)

	type ItemVO struct {
		SkuId        int     `json:"skuid"`
		ProductId    int     `json:"productid"`
		ProductName  string  `json:"productname"`
		ProductPrice float64 `json:"productprice"`
		ProductType  string  `json:"productType"`
		ProductNum   int     `json:"productnum"`
		ProductImg   string  `json:"productimg"`
	}

	// map: shop_name -> []ItemVO
	shopItemsMap := make(map[string][]ItemVO)

	// 将购物车项按店铺分组
	for _, item := range cartItems {
		priceDecimal, _ := decimal.NewFromString(item.Price)
		price, _ := priceDecimal.Float64()
		itemVo := ItemVO{
			SkuId:        item.SkuId,
			ProductId:    item.ProductId,
			ProductName:  item.Name,
			ProductPrice: price,
			ProductType:  item.GetBriefAttr(),
			ProductNum:   item.Count,
			ProductImg:   item.Image,
		}
		shopItemsMap[item.ShopName] = append(shopItemsMap[item.ShopName], itemVo)
	}

	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"msg":     "查询成功",
		"data":    shopItemsMap,
	})
}
