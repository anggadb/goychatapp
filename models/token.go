package models

import (
	"goychatapp/lib"
	"time"
)

type Token struct {
	ID        int64     `json:"id"`
	Token     string    `json:"token"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UserId    int64     `json:"user_id"`
	Active    bool      `json:"active"`
}

func GetToken(token, tokenType string, userID int64) error {
	db := lib.CreateConnection()
	defer db.Close()
	query := "SELECT * FROM tokens WHERE user_id=$1 AND token=$2 AND type=$3"
	err := db.QueryRow(query, userID, token, tokenType)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}
