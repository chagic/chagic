package model

type Chat struct {
	Base

	Name             string   `json:"name"`
	Type             ChatType `json:"type"`
	ParticipantCount int      `json:"participant_count"`
	ActiveUsers      string   `json:"active_users"`
}

type ChatType string

const (
	ChatTypePrivate ChatType = "private"
	ChatTypeGroup   ChatType = "group"
)
