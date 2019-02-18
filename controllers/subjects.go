package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SubjectsController struct {
	beego.Controller
}

func (subCtr *SubjectsController) Get() {
	o := orm.NewOrm()
	var maps []orm.Params
	_, err := o.QueryTable("mp_english_subjects").Values(&maps)
	if err != nil {
		fmt.Println(err)
		return
	}
	subCtr.Data["json"] = &maps
	subCtr.ServeJSON()
}
