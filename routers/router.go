package routers

import (
	"web-app-test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	profile := beego.NewNamespace("/profile",
		beego.NSRouter("", &controllers.ProfileController{}, "get:ListProfile"),
		beego.NSRouter("/add", &controllers.ProfileController{}, "post:AddProfile"),
		beego.NSRouter("/:id/addHobby", &controllers.ProfileController{}, "post:AddHobby"),
	)
	beego.AddNamespace(profile)
}
