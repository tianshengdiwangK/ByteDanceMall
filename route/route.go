package route

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/controller/shopAdmin"
	"login_register_demo/controller/user"
	settings "login_register_demo/utils/setting"
)

func Init_route() {
	gin.SetMode(settings.ServerSetting.RunMode)
	router := gin.Default()
	//验证器注册

	//router.Use(middleware.JwtToken())

	userOp := router.Group("/user")
	{
		userOp.POST("/login", user.UserLoginT)
		userOp.POST("/register", user.UserRegisterT)

	}

	shopAdminOp := router.Group("/shop/admin")
	{
		shopAdminOp.POST("/product/image/add", shopAdmin.UploadProductImage)
		shopAdminOp.POST("/product/add", shopAdmin.AddProduct)
		shopAdminOp.POST("/product/sku/add", shopAdmin.ProductAddSku)
	}

	//shopOp := router.Group("/mall")
	//{
	//	shopOp.GET("/classification", mall.GetMallCategory)
	//}

	//cartOp := router.Group("/api/user/cart")
	//{
	//	cartOp.GET("/all", cart.GetCartAll)
	//	cartOp.POST("/add", cart.CartAddProduct)
	//}

	router.Run(settings.ServerSetting.HttpPort)

}
