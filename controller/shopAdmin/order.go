package shopAdmin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"login_register_demo/model"
	"net/http"
)

func GetOrderList(c *gin.Context) {
	shopId, ok := c.GetQuery("shop_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"msg ":    "请求参数错误",
		})
	}
	res := make([]model.Order, 0)
	err := config.Engine.Table("order").
		Where("shop_id = ?", shopId).Find(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"msg ":    "内部查询错误",
		})
	}
	c.JSON(http.StatusOK, res)
}

func GetOrderItem(c *gin.Context) {
	body := struct {
		User_address model.UserAddress
		Items	 []model.OrderItem
	}{}
	order_id := c.Query("order_id")
	fmt.Println(order_id)
	order := new(model.Order)
	_, err1 := config.Engine.Table("order").Where("order_id=?", order_id).Get(order)
	if err1 != nil {
		fmt.Println(err1)
	}
	address := new(model.UserAddress)

	_, err2 := config.Engine.Table("user_address").Where("user_id = ?", order.UserId).Get(address)
	body.User_address = *address
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"msg ":    "内部查询用户地址错误",
		})
	}
	items := make([]model.OrderItem, 0)
	err3 := config.Engine.Table("order_item").Where("order_id = ?", order.OrderId).Find(&items)
	body.Items = items
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"msg ":    "内部查询订单项错误",
		})
	}

	c.JSON(http.StatusOK, body)
}