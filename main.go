package main

import (
	"login_register_demo/config"
	"login_register_demo/route"
	settings "login_register_demo/utils/setting"
)

func main() {
	settings.Init_setting()
	config.InitDBXorm()
	route.Init_route()
}
