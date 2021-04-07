package models

import (
	"fmt"
	"goychatapp/lib"
	"time"

	"github.com/lib/pq"
)

type Room struct {
	ID           uint         `json:"id"`
	SenderId     uint         `json:"sender_id"`
	Participants []string     `json:"participants"`
	Type         string       `json:"type"`
	CreatedAt    time.Time    `json:"created_at"`
	DeletedAt    lib.NullTime `json:"deleted_at"`
}

func CreateRoom(r Room) (int, error) {
	var roomId int
	db := lib.CreateConnection()
	defer db.Close()
	query := "INSERT INTO rooms (sender_id,participants,type) VALUES ($1,$2,$3) RETURNING id"
	err := db.QueryRow(query, r.SenderId, pq.Array(r.Participants), r.Type).Scan(&roomId)
	if err != nil {
		return roomId, err
	}
	return roomId, nil
}
func GetRoom(id string) (Room, error) {
	var room Room
	db := lib.CreateConnection()
	defer db.Close()
	query := "SELECT * FROM rooms WHERE id=$1 AND deleted_at IS NULL"
	err := db.QueryRow(query, id).Scan(&room.ID, &room.SenderId, pq.Array(&room.Participants), &room.Type, &room.CreatedAt, &room.DeletedAt)
	if err != nil {
		return room, err
	}
	return room, nil
}
func GetAllRooms(r Room, orderBy, order string, page, perPage int) ([]Room, error) {
	var rooms []Room
	db := lib.CreateConnection()
	defer db.Close()
	limit := perPage
	if limit == 0 {
		limit = 1000000
	}
	offset := limit * (page - 1)
	condition, err := lib.DynamicFilters(r)
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
		var rm Room
		err = rows.Scan(&rm.ID, &rm.SenderId, &rm.Participants, &rm.Type, &rm.CreatedAt, &rm.DeletedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, rm)
	}
	return rooms, nil
}
