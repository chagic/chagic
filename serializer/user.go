package serializer

import "chagic/model"

type TokenData struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}
