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
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Photo     lib.NullString `json:"photo"`
	Active    bool           `json:"active"`
	Verified  bool           `json:"verified"`
	CreatedAt time.Time      `json:"created_at"`
}

// func CreateUser(user User) *User {
// 	db := lib.CreateConnection()
// 	defer db.Close()

// }
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
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Password, &user.Photo, &user.Active, &user.Verified, &user.CreatedAt)
		if err != nil {
			log.Fatalf("Error while fetch datas. %v", err)
		}
		users = append(users, user)
	}
	return users, err
}
