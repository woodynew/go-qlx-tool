package routers

import (
	"go-qlx-tool/controllers"

	beego "github.com/beego/beego/v2/server/web"
	context "github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/qulaxin", &controllers.MainController{}, "get:Qulaxin")
	beego.Router("/error", &controllers.MainController{}, "get:Error")

	beego.Get("/test", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})

	ns := beego.NewNamespace("/v1",
		beego.NSGet("/test", func(ctx *context.Context) {
			ctx.Output.Body([]byte("hello world"))
		}),
		// beego.NSNamespace("/export",
		// 	beego.NSRouter("/test", &controllers.ExportController{}, "get:exportTest"),
		// ),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(beego.NewNamespace("/export",
		beego.NSRouter("/test", &controllers.ExportController{}, "get:ExportTest"),

		beego.NSRouter("/get-suning-b2", &controllers.ExportController{}, "get:GetExportSuningB2"),
		beego.NSRouter("/export-suning-b2", &controllers.ExportController{}, "get:ExportSuningB2"),
	))

}
