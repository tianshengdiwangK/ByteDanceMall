package shopAdmin

import (
	"github.com/gin-gonic/gin"
	"log"
	"login_register_demo/config"
	"net/http"
)

func GetCategoryId(c *gin.Context) {
	categoryId := c.Request.URL.Query().Get("goods_attribute_category_id")
	query, err := config.Engine.Query("select * from goods_category where category_id = ?", categoryId)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"goods_attribute_category_id": categoryId,
		"query":                       query,
	})
}
