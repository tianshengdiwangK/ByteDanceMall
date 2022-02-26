package shopAdmin

import (
	"github.com/gin-gonic/gin"
	"log"
	"login_register_demo/config"
	"login_register_demo/model"
	"login_register_demo/utils/errmsg"
	"regexp"
)

// ProductAddSku 添加SKU接口
// 	POST /shop/admin/product/sku/add
func ProductAddSku(c *gin.Context) {
	body := struct {
		GoodsId int              `json:"goods_id" binding:"required"`
		Skus    []model.GoodsSku `json:"skus" binding:"required,dive"`
	}{}

	// 参数绑定与校验
	if err := c.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		code := errmsg.ERROR_JSON_TYPE_WRONG
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
		return
	}

	// 验证goods_id是否存在
	goods := &model.Goods{Id: body.GoodsId}
	exist, err := config.Engine.Exist(goods)
	if err != nil {
		log.Println(err.Error())
		code := errmsg.ERROR_DATABASE_WRONG
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
		return
	}
	if !exist { // goods_id不存在
		code := errmsg.ERROR_PRODUCT_NOT_EXIST
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
		return
	}

	// 校验skus.price字段
	for index, sku := range body.Skus {
		matched, _ := regexp.MatchString(`^([1-9]\d{0,7}|0)(\.\d{1,2})?$`, sku.Price)
		if !matched {
			code := errmsg.ERROR_PRICE_WRONG
			c.JSON(200, gin.H{
				"success": false,
				"code":    code,
				"msg":     errmsg.GetErrMsg(code),
			})
			return
		}
		body.Skus[index].GoodsId = body.GoodsId // 设置sku的goods_id
	}

	// 插入sku
	affected, err := config.Engine.Insert(body.Skus)
	log.Printf("插入%+v条sku\n", affected)
	if err != nil {
		log.Printf("insert sku, affected %d, err:\n%+v\n", affected, err.Error())
		code := errmsg.ERROR_DATABASE_WRONG
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"msg":     "添加成功",
	})
}
