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
	beego.Router("/checkLogin", &controllers.CheckLoginController{}, "POST:CheckLogin")
	beego.Router("/logout", &controllers.LogoutController{}, "*:Logout")
	beego.Router("/password", &controllers.PasswordController{}, "*:Password")
	beego.Router("/updatePassowrd", &controllers.UpdatePasswordController{}, "*:UpdatePassword")

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
	beego.Router("/providerUpdateData", &controllers.ProviderUpdateDataController{}, "*:ProviderUpdateData")
	beego.Router("/providerDel", &controllers.ProviderDelController{}, "*:ProviderDel")

	// 账单管理
	beego.Router("/billList", &controllers.BillListController{}, "*:BillList")
	beego.Router("/billAdd", &controllers.BillAddController{}, "*:BillAdd")
	beego.Router("/billAddData", &controllers.BillAddDataController{},"*:BillAddData")
	beego.Router("/billView", &controllers.BillViewController{}, "*:BillView")
	beego.Router("/billDel", &controllers.BillDelController{}, "*:BillDel")
	beego.Router("/billUpdate", &controllers.BillUpdateController{}, "GET:BillUpdate")
	beego.Router("/billUpdateData", &controllers.BillUpdateDataController{}, "GET:BillUpdateData")
}
