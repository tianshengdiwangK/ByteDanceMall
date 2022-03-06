package shopAdmin


import (
	"fmt"
	"login_register_demo/config"
	"login_register_demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGoodsList(c *gin.Context) {
	shopId, ok := c.GetQuery("shop_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"msg ":    "请求参数错误",
		})
	}

	type queryRes struct {
		model.Cart          `xorm:"extends"`
		model.GoodsImage    `xorm:"extends"`
	}

	res := make([]queryRes, 0)
	err := config.Engine.Table("goods").
			Join("INNER", "goods_image", "goods.id = goods_image.goods_id").
			Where("shop_id = ?", shopId).Find(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"msg ":    "内部查询错误",
		})
	}

	c.JSON(http.StatusOK, res)
	fmt.Println(res);
}
