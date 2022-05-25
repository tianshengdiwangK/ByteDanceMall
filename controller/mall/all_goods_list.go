package mall
import (
	"fmt"
	"log"
	"login_register_demo/config"
	"login_register_demo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGoodsInformation(c *gin.Context) {
	type goodInformation struct {
		productid   int
		productname string
		price       float32
		img         string
	}
	res := make([][]string, 0)
	Goods := make([]model.Goods, 0)
	config.Engine.Table("goods").Where("id>?", -1).Find(&Goods)
	for _, val := range Goods {
		var cur goodInformation
		cur.productid = val.Id
		cur.productname = val.Name
		p, err := strconv.ParseFloat(val.Price, 32)
		if err != nil {
			log.Panic(err.Error())
		}
		cur.price = float32(p)
		goodsImg := new(model.GoodsImage)
		config.Engine.Table("goods_image").Where("goods_id=?", cur.productid).Get(goodsImg)
		if goodsImg.Image != "" {
			cur.img = goodsImg.Image
			curS := []string{"productid:" + strconv.Itoa(cur.productid), "productname:" + cur.productname, "price:" + fmt.Sprintf("%f", cur.price), "img:" + cur.img}
			res = append(res, curS)
		}
	}
	c.JSON(http.StatusOK, res)
}
