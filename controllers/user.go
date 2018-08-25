package controllers

import (
	"fmt"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mpEnglishBe/models"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	u.Ctx.WriteString("get user")
}

func (u *UserController) Post() {
	user := models.User{ Username: "user123", Password: "pwd123", WxOpenID: "wxopen123", UserType: "1"}
	o := orm.NewOrm()
	insertId, err := o.Insert(&user)
	if err != nil {
		fmt.Println(err)
		u.Ctx.WriteString("Some error occured")
		return
	}
	u.Ctx.WriteString("insertId:" + strconv.Itoa(int(insertId)))
}
