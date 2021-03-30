package models

import "time"

type Token struct {
	ID        int64     `json:"id"`
	Token     string    `json:"token"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UserId    int64     `json:"user_id"`
	Active    bool      `json:"active"`
}
