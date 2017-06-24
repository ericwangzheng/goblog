package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/changepass`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/changepass`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "DoAdd",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/edit/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "DoUpdate",
			Router: `/edit/:id([0-9]+)`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
