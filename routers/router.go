package routers

import (
	"blueSupermarket/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    //首页
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/password", &controllers.PasswordController{}, "*:Password")

	// 用户管理
    beego.Router("/userList", &controllers.UserListController{}, "*:UserList")
    beego.Router("/userAdd", &controllers.UserAddController{}, "*:UserAdd")
	beego.Router("/userAddData", &controllers.UserAddDataController{}, "GET:UserAddData")
	beego.Router("/userDel", &controllers.UserDelController{}, "GET:UserDel")
	beego.Router("/userUpdate", &controllers.UserUpdateController{}, "*:UserUpdate")
	beego.Router("/userUpdateData", &controllers.UserUpdateDataController{}, "*:UserUpdateData")
	beego.Router("/userView", &controllers.UserViewController{}, "*:UserView")

	// 供应商管理
	beego.Router("/providerList", &controllers.ProviderController{}, "*:ProviderList")
	beego.Router("/providerAdd", &controllers.ProviderAddController{}, "*:ProviderAdd")
	beego.Router("/providerAddData", &controllers.ProviderAddDataController{}, "*:ProviderAddData")
	beego.Router("/providerView", &controllers.ProviderViewController{}, "*:ProviderView")
	beego.Router("/providerUpdate", &controllers.ProviderUpdateController{}, "*:ProviderUpdate")
	beego.Router("/providerDel", &controllers.ProviderDelController{}, "*:ProviderDel")
}
