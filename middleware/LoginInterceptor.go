package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(config.Jwtkey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
// 生成token
func SetToken(username string) (string,int) {
	expireTime := time.Now().Add(10 * time.Hour)	//过期时间10小时
	setClaim := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),	//过期时间（时间戳）
			Issuer: "ginblog",	//发行者
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,setClaim) //生成token
	token,err := reqClaim.SignedString(JwtKey)                     //转换为字符串
	if err != nil {
		return "",403
	}
	return token,200
}
// 验证token
func CheckToken(token string) (*MyClaims,int) {
	setToken,_ := jwt.ParseWithClaims(token,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey,nil
	})
	if key,_ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key,200
	}else {
		return nil,405
	}
}
//jwt中间件
// 全局验证器
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		fmt.Println("jwt中间件正常运行----------")
		//非登录路劲进行token验证
		if strings.Index(path, "login") < 0 {
				//token := c.Request.Header.Get("token")
			token := c.Request.Header.Get("token")
			if token!=""{
				token=token[1:]
			}
				fmt.Println("token为："+token)
				if token == "" {
					c.JSON(
						http.StatusOK,
						gin.H{
							"success": false,
							"code":    405,
							"msg":     "无权访问",
						},
					)
					c.Abort()
					return
				}

			key, tCode := CheckToken(token)

			if tCode == 405{

				c.JSON(http.StatusOK,
					gin.H{
						"success": false,
						"code":    405,
						"msg":     "token错误",
					},
				)
				c.Abort()
				return
			}
			if time.Now().Unix() > key.ExpiresAt {
				c.JSON(http.StatusOK,
					gin.H{
						"success": false,
						"code":    405,
						"msg":     "超时",
					},
				)
				c.Abort()
				return
			}
			c.Set("username",key.Username)
			c.Next()
		}
	}
}
//局部验证器 检验用户权限，普通用户无法访问admin后台
func CheckAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println("checkAdminAuth。。。。。。。。。。。。。。。。")
		token := c.Request.Header.Get("token")
		state:=token[0]-'0'
		token=token[1:]
		if state==2{
			c.JSON(http.StatusOK,
				gin.H{
					"success": false,
					"code":    405,
					"msg":     "没有权限",
				},
			)
			c.Abort()
			return
		}
		c.Next()
	}
}