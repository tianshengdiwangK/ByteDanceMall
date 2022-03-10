package cart

import (
	"github.com/gin-gonic/gin"
	"log"
	"login_register_demo/config"
	"login_register_demo/model"
	"login_register_demo/utils/errmsg"
	"net/http"
)

type CartItemDTO struct {
	SkuId      int `json:"skuid" binding:"required"`
	ProductNum int `json:"productcount" binding:"required,gt=0"`
}

// CartAddProduct
//  POST /cart/add
// 添加购物车商品
func CartAddProduct(c *gin.Context) {
	// 从context中获取user_id
	val, _ := c.Get("user_id")
	userId := val.(int) // 断言为int

	// 参数绑定与校验
	var cartItemDTO CartItemDTO
	err := c.ShouldBindJSON(&cartItemDTO)
	if err != nil {
		code := errmsg.ERROR_JSON_TYPE_WRONG
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	// 校验sku是否存在
	goodsSku := model.GoodsSku{}
	has, err := config.Engine.Id(cartItemDTO.SkuId).Exist(&goodsSku)
	if err != nil || has == false {
		code := errmsg.ERROR_PRODUCT_NOT_EXIST
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": errmsg.GetErrMsg(code),
		})
		return
	}

	// 判断用户购物车中是否已存在该sku
	cart := model.Cart{
		UserId: userId,
		SkuId:  cartItemDTO.SkuId,
	}
	has, err = config.Engine.Get(&cart)
	if !has { // 用户购物车中不存在该sku
		cartItemInfo := struct {
			model.GoodsSku `xorm:"extends"`
			model.Goods    `xorm:"extends"`
			model.Shop     `xorm:"extends"`
		}{}
		_, err := config.Engine.Table("goods_sku").
			Join("INNER", "goods", "goods_sku.goods_id = goods.id").
			Join("INNER", "shop", "goods.shop_id = shop.shop_id").
			Where("goods_sku.id = ?", cartItemDTO.SkuId).
			Get(&cartItemInfo)
		if err != nil {
			log.Println(err)
			code := errmsg.ERROR_DATABASE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Message": errmsg.GetErrMsg(code),
			})
			return
		}

		cart = model.Cart{
			UserId:    userId,
			Name:      cartItemInfo.Name,
			Count:     cartItemDTO.ProductNum,
			ProductId: cartItemInfo.GoodsId,
			SkuId:     cartItemDTO.SkuId,
			ShopName:  cartItemInfo.ShopName,
		}
		_, err = config.Engine.Insert(&cart)
		if err != nil {
			code := errmsg.ERROR_DATABASE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Message": errmsg.GetErrMsg(code),
			})
			return
		}

		code := errmsg.SUCCSE
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": errmsg.GetErrMsg(code),
		})
	} else { // 用户购物车中已存在该sku
		cart.Count += cartItemDTO.ProductNum
		config.Engine.Id(cart.Id).Cols("count").Update(&cart)
		code := errmsg.SUCCSE
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": errmsg.GetErrMsg(code),
		})
	}
}
