package route

import (
	"login_register_demo/controller/shopAdmin"
	"login_register_demo/controller/user"
	settings "login_register_demo/utils/setting"

	"github.com/gin-gonic/gin"
)

func Init_route() {
	gin.SetMode(settings.ServerSetting.RunMode)
	router := gin.Default()


	//验证器注册
	router.Use(middleware.JwtToken())
	userOp := router.Group("/user")
	{
		userOp.GET("/login", user.UserLoginG)
		userOp.GET("/register", user.UserRegisterG)

	}

	//验证器注册


	//router.Use(middleware.JwtToken())


	shopOp := router.Group("/mall",middleware.CheckAdminAuth())

	userOp := router.Group("/user")

	{
		userOp.POST("/login", user.UserLoginT)
		userOp.POST("/register", user.UserRegisterT)

	}


	cartOp := router.Group("/api/user/cart",middleware.CheckAdminAuth())

	shopAdminOp := router.Group("/shop/admin")

	{
		shopAdminOp.POST("/product/image/add", shopAdmin.UploadProductImage)
		shopAdminOp.POST("/product/add", shopAdmin.AddProduct)
		shopAdminOp.POST("/product/sku/add", shopAdmin.ProductAddSku)
		shopAdminOp.GET("/attribute_category", shopAdmin.GetAttributeCategory)
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
