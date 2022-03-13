package shopAdmin

import (
	"login_register_demo/config"
	"login_register_demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetShopInfo(c *gin.Context) {
	shopId, ok := c.GetQuery("shop_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"msg ":    "请求参数错误",
		})
	}

	res := make([]model.Shop, 0)
	err := config.Engine.Table("shop").Where("shop_id = ?", shopId).Find(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"msg ":    "内部查询错误",
		})
	}

	//封装结果
	type queryRes struct {
		shopName string `json:"shopname"`
		shopInfo string `json:"shopinfo"`
	}
	r := make([]queryRes, len(res))
	for i, val := range res {
		r[i].shopName = val.ShopName
		r[i].shopInfo = val.ShopDescription
	}

	c.JSON(http.StatusOK, r)

}
