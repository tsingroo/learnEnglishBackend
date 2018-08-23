package models

// Lesson is a Model struct
type Lesson struct {
	ID          int
	LesssonName string
	Memo        string
	CreatorID   int
}

// TableName set the right db mapping table
func (l *Lesson) TableName() string {
	return "mp_english_lessons"
}
