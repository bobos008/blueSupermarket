package routers

import (
	"blueSupermarket/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.UserController{}, "*:Index")
    beego.Router("/userList", &controllers.UserListController{}, "*:UserList")
    beego.Router("/userAdd", &controllers.UserAddController{}, "*:UserAdd")
	beego.Router("/userAddData", &controllers.UserAddDataController{}, "GET:UserAddData")
	beego.Router("/userUpdate", &controllers.UserUpdateController{}, "*:UserUpdate")
	beego.Router("/userView", &controllers.UserViewController{}, "*:UserView")
}
