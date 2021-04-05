package models

import (
	"fmt"
	"goychatapp/lib"
	"time"
)

type Files struct {
	ID        uint      `json:"id"`
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
func GetFile(id string) (Files, error) {
	var f Files
	db := lib.CreateConnection()
	defer db.Close()
	query := "SELECT * FROM files WHERE id=$1"
	err := db.QueryRow(query, id).Scan(&f.ID, &f.UserId, &f.Name, &f.Type, &f.CreatedAt, &f.Path)
	if err != nil {
		return f, err
	}
	return f, nil
}
func GetAllFiles(file Files, orderBy, order string, page, perPage int) ([]Files, error) {
	var files []Files
	db := lib.CreateConnection()
	defer db.Close()
	limit := perPage
	if limit == 0 {
		limit = 1000000
	}
	offset := limit * (page - 1)
	condition, err := lib.DynamicFilters(file)
	if err != nil {
		return nil, err
	}
	pagination := fmt.Sprintf("ORDER BY %s %s LIMIT %d OFFSET %d", orderBy, order, limit, offset)
	query := "SELECT * FROM files " + condition + pagination
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var f Files
		err = rows.Scan(&f.ID, &f.UserId, &f.Name, &f.Type, &f.CreatedAt, &f.Path)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	return files, nil
}
func DeleteFile(id string) error {
	db := lib.CreateConnection()
	defer db.Close()
	query := "DELETE FROM files WHERE id=$1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
