package route

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/middleware"

	"login_register_demo/controller/cart"
	"login_register_demo/controller/mall"
	"login_register_demo/controller/user"
)

func Init_route() {
	router := gin.Default()

	//验证器注册
	router.Use(middleware.JwtToken())
	userOp := router.Group("/user")
	{
		userOp.GET("/login", user.UserLoginG)
		userOp.GET("/register", user.UserRegisterG)

	}


	shopOp := router.Group("/mall",middleware.CheckAdminAuth())
	{
		shopOp.GET("/classification", mall.GetMallCategory)
	}


	cartOp := router.Group("/api/user/cart",middleware.CheckAdminAuth())
	{
		cartOp.GET("/all", cart.GetCartAll)
	}




	router.Run(":8080")
}
