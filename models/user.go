package models

// User is a Model struct
type User struct {
	ID       int
	Username string
	Password string
	WxOpenID string
	UserType string
}

// TableName set the right db mapping table
func (u *User) TableName() string {
	return "mp_english_users"
}
