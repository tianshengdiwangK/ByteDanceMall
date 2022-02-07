package route

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/controller/user"
)

func Init_route() {
	router := gin.Default()
	userOp := router.Group("/user")
	{
		userOp.GET("/login", user.UserLoginG)
		userOp.GET("/register", user.UserRegisterG)

		userOp.GET("/classification", user.GetMallCategory)
	}
	router.Run(":8080")
}
