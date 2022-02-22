package config

import (
	"fmt"
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

