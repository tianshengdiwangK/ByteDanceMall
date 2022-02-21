package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	settings "login_register_demo/utils/setting"
	"strings"
	"xorm.io/core"
)

// 连接数据库参数

//var DB *sql.DB
var Engine *xorm.Engine


// 设置jwt的key
var Jwtkey = "4ti7ng2y0u"
=======




// 连接数据库
func InitDBXorm() {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	var userName = settings.DBSetting.Username
	var password = settings.DBSetting.Password
	var ip = settings.DBSetting.Host
	var port = settings.DBSetting.Port
	var dbName = settings.DBSetting.DBName
	var err error
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	Engine, err = xorm.NewEngine("mysql", path)
	if err != nil {
		fmt.Println("创建engine失败")
		return
	}
	Engine.ShowSQL(true)
	Engine.SetTableMapper(core.SnakeMapper{})
	if err := Engine.Ping(); err != nil {
		fmt.Println("连接数据库失败")
		return
	}
	fmt.Println("连接数据库成功")
}

//初始化redis连接池
func InitRedisPool(){
	Pool = &redis.Pool{     //实例化一个连接池
		MaxIdle:16,    //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:0,    //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout:300,    //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn ,error){     //要连接的redis数据库
			return redis.Dial("tcp","localhost:6379")
		},
	}
}