package route

import (
	"login_register_demo/controller/shopAdmin"
	"login_register_demo/controller/user"
	"login_register_demo/middleware"
	settings "login_register_demo/utils/setting"

	"github.com/gin-gonic/gin"
)

func Init_route() {
	gin.SetMode(settings.ServerSetting.RunMode)
	router := gin.Default()
	//验证器注册


	userOp := router.Group("/user")
	{
		userOp.POST("/login", user.UserLoginT)
		userOp.POST("/register", user.UserRegisterT)
	}
	router.Use(middleware.JwtToken())
	shopAdminOp := router.Group("/shop/admin",middleware.CheckAdminAuth()) //商家后台操作，普通用户不可操作，需要权限认证是否为商家
	{
		//商品操作
		shopAdminOp.GET("/category",shopAdmin.GetShopCategoryName)//得到店铺中所有分类
		shopAdminOp.POST("/category/icon/add",shopAdmin.UploadCategoryIcon)//向店铺中插入分类图标
		shopAdminOp.POST("/category/add",shopAdmin.InsertShopCategory)//向店铺中插入分类
		shopAdminOp.GET("/attribute_category", shopAdmin.GetAttributeCategory) //得到店铺商品类型
		shopAdminOp.GET("/product/list", shopAdmin.GetGoodsList)	//得到店铺商品列表
		shopAdminOp.POST("/product/image/add", shopAdmin.UploadProductImage)	//上传商品图片
		shopAdminOp.POST("/product/add", shopAdmin.AddProduct) 	//添加商品
		shopAdminOp.POST("/product/sku/add", shopAdmin.ProductAddSku) //添加sku
		shopAdminOp.POST("attribute_category/add", shopAdmin.AddAttributeCategory) //添加商品类型
		shopAdminOp.POST("attribute/add", shopAdmin.AddAttribute) //添加商品类型中的具体规格

		//订单操作
		shopAdminOp.GET("order/list",shopAdmin.GetOrderList) //得到订单列表
		shopAdminOp.GET("order/list/detail",shopAdmin.GetOrderItem) //得到订单详情
	}

	router.Run(settings.ServerSetting.HttpPort)

}
