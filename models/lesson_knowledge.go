package models

// LessonKnowledge is a Model struct
type LessonKnowledge struct {
	ID       int
	VoiceURL string
	ImageURL string
	LessonID int
}

// TableName set the right db mapping table
func (lk *LessonKnowledge) TableName() string {
	return "mp_english_less_knowledge"
}
