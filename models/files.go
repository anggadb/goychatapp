package models

import (
	"goychatapp/lib"
	"time"
)

type Files struct {
	ID        int64     `json:"id"`
	UserId    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	Path      string    `json:"path"`
}

func CreateFile(file Files) (string, error) {
	var path string
	db := lib.CreateConnection()
	defer db.Close()
	query := "INSERT INTO files (user_id,name,type,path) VALUES ($1,$2,$3,$4) RETURNING path"
	err := db.QueryRow(query, file.UserId, file.Name, file.Type, file.Path).Scan(&path)
	if err != nil {
		return err.Error(), err
	}
	return path, nil
}
