package models

import (
	"context"
	"goychatapp/lib"
	"log"
	"time"
)

type User struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Username  lib.NullString `json:"username"`
	Password  string         `json:"password"`
	Photo     lib.NullString `json:"photo"`
	Active    bool           `json:"active"`
	Verified  bool           `json:"verified"`
	CreatedAt time.Time      `json:"created_at"`
	Type      string         `json:"type"`
}

func CreateUser(user User) (string, error) {
	var id int64
	var token string
	db := lib.CreateConnection()
	defer db.Close()
	randomString := lib.RandomStrings(20)
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	query := "INSERT INTO users (name,password,email) VALUES ($1,$2,$3) RETURNING id"
	err = tx.QueryRowContext(ctx, query, user.Name, user.Password, user.Email).Scan(&id)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	query = "INSERT INTO tokens (token,type,user_id) VALUES ($1,$2,$3) RETURNING token"
	err = tx.QueryRowContext(ctx, query, randomString, "verify-account", id).Scan(&token)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return token, nil
}
func GetUser(email string) (User, error) {
	var u User
	var e error = nil
	db := lib.CreateConnection()
	defer db.Close()
	query := "SELECT * FROM users WHERE email=$1"
	row := db.QueryRow(query, email)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Username, &u.Password, &u.Photo, &u.Active, &u.Verified, &u.CreatedAt, &u.Type)
	if err != nil {
		e = err
	}
	return u, e
}
func GetAllUsers() ([]User, error) {
	db := lib.CreateConnection()
	defer db.Close()
	var users []User
	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error execute query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err = rows.Scan(&user)
		if err != nil {
			log.Fatalf("Error while fetch datas. %v", err)
		}
		users = append(users, user)
	}
	return users, err
}
func UpdateUser(email string, user User) (int64, error) {
	var e error = nil
	db := lib.CreateConnection()
	defer db.Close()
	query := "UPDATE users SET name=$1, email=$2, username=$3, photo=$4 WHERE email=$5"
	res, err := db.Exec(query, user.Name, user.Email, user.Username, user.Password, user.Photo, email)
	if err != nil {
		e = err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		e = err
	}
	return rows, e
}
func DeleteUser(email string) (int64, error) {
	db := lib.CreateConnection()
	defer db.Close()
	query := "DELETE FROM users WHERE email=$1"
	res, err := db.Exec(query, email)
	if err != nil {
		return 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, err
}
