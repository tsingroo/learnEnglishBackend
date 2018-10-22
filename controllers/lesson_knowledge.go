package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LessonKnowledgeController struct {
	beego.Controller
}

func (lk *LessonKnowledgeController) Get() {
	lesson_id, _ := lk.GetInt("lesson_id")
	// lk.Ctx.WriteString("get lesson Knowledge")
	o := orm.NewOrm()
	var maps []orm.Params
	o.QueryTable("mp_english_less_knowledge").Filter("lesson_id", lesson_id).Values(&maps)
	lk.Data["json"] = &maps
	lk.ServeJSON()
}
