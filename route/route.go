package route

import (
	"github.com/gin-gonic/gin"

	"login_register_demo/controller/cart"
	"login_register_demo/controller/mall"
	"login_register_demo/controller/user"
	settings "login_register_demo/utils/setting"
)

func Init_route() {
	gin.SetMode(settings.ServerSetting.RunMode)
	router := gin.Default()
	userOp := router.Group("/user")
	{
		userOp.GET("/login", user.UserLoginG)
		userOp.GET("/register", user.UserRegisterG)

	}

<<<<<<< HEAD

=======
>>>>>>> wugang
	shopOp := router.Group("/mall")
	{
		shopOp.GET("/classification", mall.GetMallCategory)
	}

<<<<<<< HEAD

	cartOp := router.Group("/api/user/cart")
	{
		cartOp.GET("/all", cart.GetCartAll)
	}



=======
	router.Run(settings.ServerSetting.HttpPort)
>>>>>>> wugang
}
