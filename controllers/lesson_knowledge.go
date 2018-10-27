package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// LessonKnowledgeController Definition
type LessonKnowledgeController struct {
	beego.Controller
}

// Get Request
func (lk *LessonKnowledgeController) Get() {
	lessonID, _ := lk.GetInt("lesson_id")
	o := orm.NewOrm()
	var maps []orm.Params
	o.QueryTable("mp_english_less_knowledge").Filter("lesson_id", lessonID).Values(&maps)
	lk.Data["json"] = &maps
	lk.ServeJSON()
}
