package shopAdmin

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"login_register_demo/model"
	"login_register_demo/utils/errmsg"
	"net/http"
)

func GetBrandList(c *gin.Context) {
	data := []model.GoodsBrand{}
	config.Engine.Find(&data)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": errmsg.GetErrMsg(code),
		"Data":    data,
	})
}
