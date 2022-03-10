package cart

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"login_register_demo/model"
	"login_register_demo/utils/errmsg"
	"net/http"
)

// CartEditProduct
//  POST /cart/sync
// 添加购物车商品
func CartEditProduct(c *gin.Context) {
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

	// 判断用户购物车中是否已存在该sku
	cart := model.Cart{
		UserId: userId,
		SkuId:  cartItemDTO.SkuId,
	}
	has, err := config.Engine.Get(&cart)
	if !has { // 用户购物车中不存在该sku
		code := errmsg.ERROR_PRODUCT_NOT_EXIST
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": errmsg.GetErrMsg(code),
		})
		return
	}
	// 用户购物车中已存在该sku
	cart.Count = cartItemDTO.ProductNum
	config.Engine.Id(cart.Id).Cols("count").Update(&cart)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": errmsg.GetErrMsg(code),
	})

}
