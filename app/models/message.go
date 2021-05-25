package models

import (
	"github.com/kamva/mgm/v3"
)

type Message struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	UserId           string `json:"user_id" bson:"user_id"`
	RoomId           int    `json:"room_id" bson:"room_id"`
	Text             string `json:"text" bson:"text"`
}

func NewMessage(userId string, roomId int, text string) *Message {
	return &Message{
		UserId: userId,
		RoomId: roomId,
		Text:   text,
	}
}
