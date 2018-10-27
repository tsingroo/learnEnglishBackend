package models

// Subject Model Struct
type Subjects struct {
	ID           int `orm:"column(id)"`
	Description  string
	ImageURL     string `orm:"column(image_url)"`
	Memo         string
	DisplayOrder int
	CreatorID    string `orm:"column(creator_id)"`
}

func (sub *Subjects) TableName() string {
	return "mp_english_subjects"
}
