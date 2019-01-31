package routers

import (
	"Publish/controllers"
	"github.com/astaxie/beego"
)


func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    beego.Router("/index",&controllers.ArticleController{},"get:ShowIndex")
    beego.Router("/add",&controllers.ArticleController{},"get:ShowAdd;post:HandleAdd")
    beego.Router("/content",&controllers.ArticleController{},"get:ShowContent")
    beego.Router("/update",&controllers.ArticleController{},"get:ShowUpdate;post:HandleUpdate")
    beego.Router("/delete",&controllers.ArticleController{},"get:HandleDelete")
    beego.Router("/addType",&controllers.ArticleController{},"get:ShowAddType")



}
