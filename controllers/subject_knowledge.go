package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SubjectKnowledgeController struct {
	beego.Controller
}

func (sk *SubjectKnowledgeController) Get() {
	var maps []orm.Params
	subID, _ := sk.GetInt("sub_id")
	o := orm.NewOrm()
	_, err := o.QueryTable("mp_english_subject_knowledge").Filter("subject_id", subID).Values(&maps)
	if err != nil {
		fmt.Println(err)
		return
	}
	sk.Data["json"] = &maps
	sk.ServeJSON()
}
