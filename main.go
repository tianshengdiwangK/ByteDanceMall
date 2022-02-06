package main

import (
	"login_register_demo/config"
	"login_register_demo/route"
)

func main()  {
	config.InitDBXorm()
	route.Init_route()
}