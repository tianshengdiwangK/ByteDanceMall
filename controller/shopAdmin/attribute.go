package shopAdmin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"login_register_demo/config"
	"login_register_demo/model"
	"login_register_demo/utils/errmsg"
)
type receivedAttribute struct {
	Name string `json:"name"`
	Select_type int `json:"select_type"`
	Goods_attribute_category string `json:"goods_attribute_category"`
	Input_list string `json:"input_list"`

}

func AddAttribute(c *gin.Context) {
	receive := receivedAttribute{}

	//参数绑定
	if err := c.ShouldBind(&receive); err != nil {
		log.Println(err.Error())
		code := errmsg.ERROR_JSON_TYPE_WRONG
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
		return
	}
	goods_attribute_category_id := queryIdByName(receive.Goods_attribute_category)
	attribute := model.GoodsAttribute{
		GoodsAttributeCategoryId: goods_attribute_category_id,
		Name: receive.Name,
		SelectType: receive.Select_type,
		InputList: receive.Input_list,
	}
	affected, err := config.Engine.Insert(attribute)
	if err != nil {
		fmt.Println(err)
	}
	if affected != 1 {
		c.JSON(200, gin.H{
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "插入商品类型成功",
		})
	}
}
func queryIdByName(name string) int {
	st2 := new(model.GoodsAttributeCategory)
	_, err := config.Engine.Where("goods_attribute_name=?", name).Get(st2)
	if err != nil {
		panic(err.Error())
	} else {
		return st2.Id
	}
}
