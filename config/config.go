package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strings"
	"xorm.io/core"
)

// 连接数据库参数
const (
	userName = "root"
	password = "123456"
	ip       = "110.42.132.130"
	port     = "3306"
	dbName   = "mall"
)

//var DB *sql.DB
var Engine *xorm.Engine

//创建redis连接池
var Pool *redis.Pool

// 连接数据库
func InitDBXorm() {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
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