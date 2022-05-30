package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

//判断是否登录
func (this *BaseController) Prepare() {
	loginuser := this.GetSession("loginuser")
	fmt.Println("loginuser---->", loginuser)
	if loginuser != nil {
		this.IsLogin = true
		this.Loginuser = loginuser
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}
