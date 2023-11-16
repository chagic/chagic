package model

type Participant struct {
	Base

	UserID int  `gorm:"column:user_id"`
	User   User `gorm:"foreignKey:UserID"`

	ChatID int  `gorm:"column:chat_id"`
	Chat   Chat `gorm:"foreignKey:ChatID"`

	Role Role `gorm:"column:role"`
}

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)
