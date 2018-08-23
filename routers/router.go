package routers

import (
	"mpEnglishBe/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/lesson", &controllers.LessonController{})
	beego.Router("/lessonknowledge", &controllers.LessonKnowledgeController{})
}
