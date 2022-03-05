package route

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/controller/shopAdmin"

	"login_register_demo/controller/cart"
	"login_register_demo/controller/mall"
	"login_register_demo/controller/user"
)

func Init_route() {
	router := gin.Default()
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
	}

	shopAdminOp := router.Group("/shop/admin")
	{
		shopAdminOp.GET("/attribute", shopAdmin.GetCategoryId)
	}

	router.Run(":8080")
}
