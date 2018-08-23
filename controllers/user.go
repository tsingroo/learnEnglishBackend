package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	u.Ctx.WriteString("get user")
}
