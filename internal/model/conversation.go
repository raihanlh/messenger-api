package model

type Conversation struct {
	Model
	Users []*User `gorm:"many2many:user_participants;"`
}
