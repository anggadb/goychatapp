package models

import (
	"goychatapp/lib"
)

type Chat struct {
	ID        uint           `json:"id"`
	RoomId    uint           `json:"room_id"`
	Text      lib.NullString `json:"text"`
	FileId    uint           `json:"file_id"`
	SenderId  uint           `json:"sender_id"`
	Read      bool           `json:"read"`
	CreatedAt lib.NullTime   `json:"created_at"`
	UpdatedAt lib.NullTime   `json:"updated_at"`
	DeletedAt lib.NullTime   `json:"deleted_at"`
}
