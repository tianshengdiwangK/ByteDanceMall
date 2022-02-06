package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strings"
	"xorm.io/core"
)

// 连接数据库参数
const (
	userName = "root"
	password = "root"
	ip       = "localhost"
	port     = "3306"
	dbName   = "users"
)

//var DB *sql.DB
var Engine *xorm.Engine

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
