package route

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/middleware"



	"login_register_demo/controller/shopAdmin"

	"login_register_demo/controller/user"
	settings "login_register_demo/utils/setting"
)

func Init_route() {
	gin.SetMode(settings.ServerSetting.RunMode)
	router := gin.Default()
	//验证器注册
	router.Use(middleware.JwtToken())

	userOp := router.Group("/user")

	{
		userOp.POST("/login", user.UserLoginT)
		userOp.POST("/register", user.UserRegisterT)

	}

	shopAdminOp := router.Group("/shop/admin",middleware.CheckAdminAuth()) //商家后台操作，普通用户不可操作，需要权限认证是否为商家

	{
		shopAdminOp.POST("/product/image/add", shopAdmin.UploadProductImage)
		shopAdminOp.POST("/product/add", shopAdmin.AddProduct)
		shopAdminOp.POST("/product/sku/add", shopAdmin.ProductAddSku)
		shopAdminOp.GET("/product/list", shopAdmin.GetGoodsList)
	}

	

	router.Run(settings.ServerSetting.HttpPort)

}
