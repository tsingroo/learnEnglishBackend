package models

// User is a Model struct
type User struct {
	ID       int    `orm:"column(id)"`
	Username string
	Password string
	WxOpenID string `orm:"column(wx_open_id)"`
	UserType string
}

// TableName set the right db mapping table
func (u *User) TableName() string {
	return "mp_english_users"
}
