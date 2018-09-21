package models

// Lesson is a Model struct
type Lesson struct {
	ID          int     `orm:"column(id)"`
	LessonName  string
	ImageUrl    string
	Memo        string
	CreatorID   int     `orm:"column(creator_id)"`
}

// TableName set the right db mapping table
func (l *Lesson) TableName() string {
	return "mp_english_lessons"
}
