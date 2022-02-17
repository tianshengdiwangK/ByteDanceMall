package route

import (
	"github.com/gin-gonic/gin"

	"login_register_demo/controller/cart"
	"login_register_demo/controller/mall"
	"login_register_demo/controller/user"
	"login_register_demo/middleware"
	settings "login_register_demo/utils/setting"
)

func Init_route() {
	gin.SetMode(settings.ServerSetting.RunMode)
	router := gin.Default()
	//验证器注册
	router.Use(middleware.TokenMiddle())
	userOp := router.Group("/user")
	{
		userOp.GET("/login", user.UserLoginG)
		userOp.GET("/register", user.UserRegisterG)

	}

	shopOp := router.Group("/mall")
	{
		shopOp.GET("/classification", mall.GetMallCategory)
	}

	cartOp := router.Group("/api/user/cart")
	{
		cartOp.GET("/all", cart.GetCartAll)
		cartOp.POST("/add", cart.CartAddProduct)
	}

	router.Run(settings.ServerSetting.HttpPort)

}
