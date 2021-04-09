package main

import (
	"fmt"
	"go-qlx-tool/models"
	_ "go-qlx-tool/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	fmt.Println("init-models")

	// set default database
	orm.RegisterDataBase("default", "mysql", "woody_gs:hsI7ksotjb^8tBXf7f7w@tcp(rm-bp12dk41hl53tvwv7xo.mysql.rds.aliyuncs.com:3306)/lezhuan_new?charset=utf8mb4&loc=Local")

	// register model
	orm.RegisterModel(new(models.ExternalImportData))
	orm.RegisterModel(new(models.ActivityChannel))

	// create table
	// orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()

}
