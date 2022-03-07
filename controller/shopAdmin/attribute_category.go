package shopAdmin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"login_register_demo/config"
	"login_register_demo/model"
	"login_register_demo/utils/errmsg"
	"net/http"
)

func GetAttributeCategory(c *gin.Context) {
	shopId, ok := c.GetQuery("shop_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"msg ":    "请求参数错误",
		})
	}

	res := make([]model.GoodsAttributeCategory, 0)
	err := config.Engine.Table("goods_attribute_category").Where("shop_id = ?", shopId).Find(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"msg ":    "内部查询错误",
		})
	}

	c.JSON(http.StatusOK, res)
}


func AddAttributeCategory(c *gin.Context) {
	var attribute_category model.GoodsAttributeCategory
	//var temp model.GoodsAttributeCategory //存放数据库查询结果
	//参数绑定
	if err := c.ShouldBind(&attribute_category); err != nil {
		log.Println(err.Error())
		code := errmsg.ERROR_JSON_TYPE_WRONG
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
		return
	}
	affected, err := config.Engine.Insert(attribute_category)
	if err != nil {
		fmt.Println(err)
	}
	if affected != 1 {
		c.JSON(200, gin.H{
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"success":  true,
			"msg":      "add success",
		})
	}
}