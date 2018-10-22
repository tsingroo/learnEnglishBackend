package models

// LessonKnowledge is a Model struct
type LessonKnowledge struct {
	ID       int     `orm:"column(id)"`
	Word     string  
	VoiceURL string  `orm:"column(voice_url)"`
	ImageURL string  `orm:"column(image_url)"`
	LessonID int     `orm:"column(lesson_id)"`
}

// TableName set the right db mapping table
func (lk *LessonKnowledge) TableName() string {
	return "mp_english_less_knowledge"
}
