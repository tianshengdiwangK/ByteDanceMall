package middleware

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"net/http"
	"strings"
	"time"
)

/*
登录拦截器
1,login访问-根据用户名生成token存入redis中并返回给客户端，并设置token生存时间为30min
2,其他访问需要验证是否携带token，并判断是否过期以及是否正确
3，正确登录时并更新token的生存时间
4，默认设置token的生存时间为30min
*/

func TokenMiddle() gin.HandlerFunc {
	return func(con *gin.Context) {
		// 根据headers 中的 token，判断用户是否登录
		path := con.Request.URL.Path
		if strings.Index(path, "login") < 0 {
			// 进行token验证
			token := con.Request.Header.Get("token")
			if token == "" {
				con.JSON(
					http.StatusOK,
					gin.H{
						"success": false,
						"code": 405,
						"msg":  "无权访问",
					},
				)
				con.Abort()
				return
			}else{
				conn := config.Pool.Get() //从连接池，取一个链接
				defer conn.Close() //函数运行结束 ，把连接放回连接池
				rep, _ := redis.String(conn.Do("Get", token))
				conn.Close()
				if rep == "" {
					con.JSON(
						http.StatusOK,
						gin.H{
							"success": false,
							"code": 405,
							"msg":  "toke无效",
						},
					)
					con.Abort()
					return
				}else{// 如果token存在的话，则更新token的生存时间,生存时间设置为30min
					conn.Do("expire",rep,30*time.Minute)
					con.Next()
				}
			}
		} else {
			con.Next()
		}
	}
}
