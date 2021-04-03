package models

import (
	"context"
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
func ActivateToken(token, tokenType string, userID int64) error {
	db := lib.CreateConnection()
	defer db.Close()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	query := "UPDATE tokens SET active=false WHERE token=$1 AND type=$2 AND user_id=$3 AND active=true"
	_, err = tx.ExecContext(ctx, query, token, tokenType, userID)
	if err != nil {
		tx.Rollback()
		return nil
	}
	query = "UPDATE users SET verified=true WHERE id=$1 AND verified=false"
	_, err = tx.ExecContext(ctx, query, userID)
	if err != nil {
		tx.Rollback()
		return nil
	}
	err = tx.Commit()
	if err != nil {
		return nil
	}
	return nil
}
