package models

type SubjectKnowledge struct {
	ID           int `orm:"column(id)"`
	Description  string
	ImageURL     string `orm:"column(image_url)"`
	Question     string
	Answer       string
	DisplayOrder int
	SubjectID    int `orm:"column(subject_id)"`
}

func (sk *SubjectKnowledge) TableName() string {
	return "mp_english_subject_knowledge"
}
