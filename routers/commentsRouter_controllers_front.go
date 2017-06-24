package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"],
		beego.ControllerComments{
			Method: "ShowArticleById",
			Router: `/articleid/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"],
		beego.ControllerComments{
			Method: "ShowArticlesByTag",
			Router: `/tag/:tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"] = append(beego.GlobalControllerRouter["github.com/nsecgo/goblog/controllers/front:FrontController"],
		beego.ControllerComments{
			Method: "Search",
			Router: `/search`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
