package shopAdmin

import (
	"fmt"
	"log"
	"login_register_demo/config"
	"login_register_demo/model"
	"login_register_demo/utils"
	"login_register_demo/utils/errmsg"
	"net/http"

	mapset "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
)

type receivedCategory struct {
	Category_Name        string `json:"category_name"`
	Category_icon        string `json:"category_icon"`
	Category_description string `json:"category_description"`
	Category_show_status int    `json:"category_show_status"`
	Parent_name          string `json:"parent_name"`
	Shop_id              int    `json:"shop_id"`
}
type GoodsCategory model.GoodsCategory

func GetShopCategoryName(c *gin.Context) {

	var shopCategoryName []interface{}
	var shopCategory []GoodsCategory
	set := mapset.NewSet()
	shopId := c.Request.URL.Query().Get("shop_id")
	err := config.Engine.Table("goods_category").Where("shop_id=?", shopId).Find(&shopCategory)
	if err != nil {
		panic(err.Error())
	} else {
		for i := 0; i < len(shopCategory); i++ {
			set.Add(shopCategory[i].CategoryName)
		}
		for val := range set.Iterator().C {
			shopCategoryName = append(shopCategoryName, val)
		}
		c.JSON(200, gin.H{
			"shop_categories": shopCategoryName,
		})
	}

}

// UploadProductImage 上传商品图片接口
//  POST /shop/admin/product/image/add multipart/form-data
// 在添加商品前调用该接口上传商品图片，单次上传一张，返回图片路径
func UploadCategoryIcon(c *gin.Context) {
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
	filepath, err := utils.SaveImage(file, "category")
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

func InsertShopCategory(c *gin.Context) {

	var receive receivedCategory
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

	parentId := queryParentId(receive.Parent_name)
	category := GoodsCategory{
		CategoryName:        receive.Category_Name,
		CategoryDescription: receive.Category_description,
		CategoryIcon:        receive.Category_icon,
		CategoryShowStatus:  receive.Category_show_status,
		ShopId:              receive.Shop_id,
		ParentId:            parentId,
	}
	affected, err := config.Engine.Insert(category)
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
			"msg":     "插入分类成功",
		})
	}
}

func queryParentId(parentName string) int {
	if parentName == "" {
		return 0
	} else {
		st2 := new(GoodsCategory)
		_, err := config.Engine.Where("category_name=?", parentName).Get(st2)
		if err != nil {
			panic(err.Error())
		} else {
			return st2.CategoryId
		}
	}
}

func GetCategoryGoods(c *gin.Context) {
	categoryId, ok := c.GetQuery("category_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"msg ":    "请求参数错误",
		})
	}

	type goodsItem struct {
		model.Goods      `xorm:"extends"`
		model.GoodsImage `xorm:"extends"`
	}
	res := make([]goodsItem, 0)
	config.Engine.Table("goods").
		Join("INNER", "goods_image", "goods_image.goods_id = cart.product_id and goods_image.is_primary = 1").
		Where("goods.goods_category_id = ?", categoryId).
		Find(&res)

	//封装结果
	type queryRes struct {
		productid   int    `json:"productid"`
		productname string `json:"productname"`
		productinfo string `json:"productinfo"`
		price       string `json:"price"`
		productimg  string `json:"productimg"`
	}
	r := make([]queryRes, len(res))
	for i, val := range res {
		r[i].productid = val.GoodsId
		r[i].productname = val.Name
		r[i].productinfo = val.Description
		r[i].price = val.Price
		r[i].productimg = val.Image
	}
	c.JSON(http.StatusOK, r)
}
