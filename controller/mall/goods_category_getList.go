package mall

import (
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"login_register_demo/utils/errmsg"
	"net/http"
)

type category struct {
	CategoryId   int
	CategoryName string
	ParentId     int
	Children     []category `xorm:"null"`
}

func GetCategoryList(c *gin.Context) {
	categoryList := make([]category, 0)
	config.Engine.Table("goods_category").Cols("category_id", "category_name", "parent_id").Find(&categoryList)
	code := errmsg.SUCCSE
	req := make([]category, len(categoryList))
	k := 0
	for i := range categoryList {
		if categoryList[i].ParentId == 0 {
			req[k].CategoryId = categoryList[i].CategoryId
			req[k].CategoryName = categoryList[i].CategoryName
			req[k].ParentId = categoryList[i].ParentId
			k++
		} else {
			for j := range req {
				if req[j].CategoryId == categoryList[i].ParentId {
					req[j].Children = append(req[j].Children, categoryList[i])
				}
			}
		}
	}
	req = req[0:k]
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": errmsg.GetErrMsg(code),
		"Data":    req,
	})
}
