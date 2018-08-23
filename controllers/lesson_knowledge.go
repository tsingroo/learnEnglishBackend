package controllers

import (
	"github.com/astaxie/beego"
)

type LessonKnowledgeController struct {
	beego.Controller
}

func (lk *LessonKnowledgeController) Get() {
	lk.Ctx.WriteString("get lesson Knowledge")
}
