package mall

import (

	"fmt"
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"login_register_demo/model"
)

func GetProductInfo(c *gin.Context) {
	product_id := c.Query("product_id")
	fmt.Println(product_id)
	//查询结果
	type Product_detail struct {
		goods model.Goods
		image []model.GoodsImage
		sku []model.GoodsSku
	}
	detail := Product_detail{}
	 _, err1 := config.Engine.Table("goods").
		Where("goods.id = ?", product_id).
		Get(&detail.goods)
	if err1!= nil {
		fmt.Println(err1)
	}
	config.Engine.Table("goods_image").
		Where("goods_image.goods_id = ?", product_id).
		Find(&detail.image)
	 config.Engine.Table("goods_sku").
		Where("goods_sku.goods_id = ?", product_id).
		Find(&detail.sku)
	//detail.sku = items

	//返回结果
	type information struct {
		id int
		name string
		des string
		price string
		category string
	}
	res := struct {
		info information
		productimg []string
		sku 		[]model.GoodsSku
	}{}

	fmt.Println(detail.goods.Name)
	res.info = information{
		id : detail.goods.Id,
		name: detail.goods.Name,
		des: detail.goods.Description,
		price: detail.goods.Price,
		category: QueryCategoryById(detail.goods.GoodsCategoryId),
	}
	fmt.Println(res.info)
	for i := 0; i < len(detail.image); i++ {
		res.productimg = append(res.productimg,detail.image[i].Image)
	}
	fmt.Println(res.productimg)

	res.sku = detail.sku

	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"msg":     "查询成功",
		"data":    res.sku,
	})
}
func QueryCategoryById(id int) string  {
	var category model.GoodsCategory
	_, err := config.Engine.Where("category_id = ?", id).Get(category)
	if err != nil {
		println(err)
	}
	return category.CategoryName
}