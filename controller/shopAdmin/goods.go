package shopAdmin

import (
	"github.com/gin-gonic/gin"
	"log"
	"login_register_demo/config"
	"login_register_demo/model"
	"login_register_demo/utils"
	"login_register_demo/utils/errmsg"
	"regexp"
)

// UploadProductImage 上传商品图片接口
//  POST /shop/admin/product/image/add multipart/form-data
// 在添加商品前调用该接口上传商品图片，单次上传一张，返回图片路径
func UploadProductImage(c *gin.Context) {
	// 从context中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		code := errmsg.ERROR_UPLOAD_WRONG
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
		return
	}

	// 保存文件，得到图片相对路径，如“/img/product/xxx.jpg”，前端访问时src为“http://static.xxx.com/img/product/xxx.jpg”
	filepath, err := utils.SaveImage(file, "product")
	if err != nil {
		log.Println(err)
		code := errmsg.ERROR_UPLOAD_WRONG
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
		"data": gin.H{
			"img": filepath,
		},
	})
}

// AddProduct 添加商品接口
// 	POST /shop/admin/product/add
// 添加商品，并添加该商品的全部图片和sku
func AddProduct(c *gin.Context) {
	body := struct {
		Info model.Goods      `json:"info" binding:"required,dive"`
		Imgs []string         `json:"imgs"`
		Skus []model.GoodsSku `json:"skus" binding:"required,dive"`
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
	log.Printf("%+v\n", body)

	// 校验skus.price字段
	for _, sku := range body.Skus {
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
	}

	// 插入product
	affected, err := config.Engine.InsertOne(&body.Info)
	log.Printf("插入%+v条商品,id：%d\n", affected, body.Info.Id)
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

	// 设置sku的goods_id
	for index := range body.Skus {
		body.Skus[index].GoodsId = body.Info.Id
	}

	// 插入sku
	affected, err = config.Engine.Insert(&body.Skus)
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

	// 插入图片
	imgs := make([]model.GoodsImage, len(body.Imgs))
	for index, img := range body.Imgs {
		imgs[index].GoodsId = body.Info.Id
		imgs[index].Image = img
	}
	affected, err = config.Engine.Insert(&imgs)
	log.Printf("插入%+v条img\n", affected)
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
