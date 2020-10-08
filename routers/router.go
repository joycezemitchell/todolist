package routers

import (
	"todolist2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/todo", &controllers.TodoController{})
	beego.Router("/generatetoken", &controllers.GeneratetokenController{})
	beego.Router("/user", &controllers.UserController{})
}
