package controllers

import (
	"github.com/astaxie/beego"
)

type LessonController struct {
	beego.Controller
}

func (l *LessonController) Get() {
	l.Ctx.WriteString("get Lesson")
}
