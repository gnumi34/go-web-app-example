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
		beego.NSRouter("/:idProfile/addHobby", &controllers.ProfileController{}, "get:AddHobby"),
		beego.NSRouter("/:idProfile/updateProfile", &controllers.ProfileController{}, "get:ShowUpdateProfile"),
		beego.NSRouter("/:idProfile/updateProfile", &controllers.ProfileController{}, "post:UpdateProfile"),
		beego.NSRouter("/:idProfile/addHobby/submit", &controllers.ProfileController{}, "post:SubmitHobby"),
		beego.NSRouter("/:idProfile/deleteProfile", &controllers.ProfileController{}, "post:DeleteProfile"),
		beego.NSRouter("/:idProfile/hobby/:idHobby/delete", &controllers.ProfileController{}, "post:DeleteHobby"),
		beego.NSRouter("/:idProfile/hobby/:idHobby/update", &controllers.ProfileController{}, "get:ShowUpdateHobby"),
		beego.NSRouter("/:idProfile/hobby/:idHobby/update", &controllers.ProfileController{}, "post:UpdateHobby"),
	)
	beego.AddNamespace(profile)
}
