package models

import (
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
	Type      string         `json:"type"`
	CreatedAt time.Time      `json:"created_at"`
}

func CreateUser(user User) int64 {
	var id int64
	db := lib.CreateConnection()
	defer db.Close()
	sql := "INSERT INTO users (name,password,email) VALUES ($1,$2,$3) RETURNING id"
	err := db.QueryRow(sql, user.Name, user.Password, user.Email).Scan(&id)
	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}
	return id
}
func GetUser(user User) (User, error) {
	var u User
	var e error = nil
	db := lib.CreateConnection()
	defer db.Close()
	sqlState := "SELECT * FROM users WHERE username = 'anggbchtr'"
	row := db.QueryRow(sqlState)
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
	sql := "SELECT * FROM users"
	rows, err := db.Query(sql)
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
