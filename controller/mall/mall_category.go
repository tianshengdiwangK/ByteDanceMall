package mall

import (
	"github.com/gin-gonic/gin"
	"log"
	"login_register_demo/config"
)

// GetMallCategory
// GET /mall/classification
// 获取平台类目列表
func GetMallCategory(c *gin.Context) {
	var categories []string
	err := config.Engine.Table("mall_category").Cols("category_name").Find(&categories)
	if err != nil {
		log.Println("get mall category:", err)
	}
	c.JSON(200, gin.H{
		"categories": categories,
	})
}
