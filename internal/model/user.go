package model

import "gitlab.com/raihanlh/messenger-api/internal/constant"

type User struct {
	Model    `swaggerignore:"true"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-" swaggerignore:"true"`
	PhotoURL string `json:"photo_url,omitempty"`
}

// Table name for gorm
func (u *User) Table() string {
	return constant.UserTable
}
