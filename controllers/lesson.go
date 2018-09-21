package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LessonController struct {
	beego.Controller
}

func (l *LessonController) Get() {
	o := orm.NewOrm()
	// 不推荐使用此种方式
	// var maps []orm.Params
	// num, err := o.Raw("SELECT * FROM mp_english_lessons").Values(&maps)
	// if err == nil {
	// 	fmt.Println(num)
	// } else {
	// 	fmt.Println(err) 
	// }
	// l.Data["json"] = &maps
	// l.ServeJSON()
	var maps []orm.Params
	_, err := o.QueryTable("mp_english_lessons").Values(&maps)
	if err != nil {
		fmt.Println(err)
		return
	}
	l.Data["json"] = &maps
	l.ServeJSON()
}
