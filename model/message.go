package model

type Message struct {
	Base
	Type   string `json:"type"`
	Msg    string `json:"msg"`
	ChatID int    `gorm:"column:chat_id" json:"chat_id"`
	Chat   Chat   `gorm:"foreignKey:ChatID" json:"-"`
	UserID int    `gorm:"column:user_id" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}
