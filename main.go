package main

import (
	_ "go-qlx-tool/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

