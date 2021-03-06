package main

import (
	"../utils"
	"database/sql"
	"time"
)

// RR typical RR structure.
type RR struct {
	ID                         int64     `json:"id"`
	PersonID                   int64     `json:"person_id"`
	ItemID                     int64     `json:"item_id"`
	PersonName                 string    `json:"reviewer_name"`
	Relationship               string    `json:"relationship"`
	RelationshipDurationInDays int       `json:"relationship_duration_in_days"`
	Rating                     int       `json:"rating"`
	Comments                   string    `json:"comments"`
	Pros                       string    `json:"pros"`
	Cons                       string    `json:"cons"`
	RelationshipDate           time.Time `json:"relationship_date"`
	HasResponse                bool      `json:"has_response"`
	IsResponse                 bool      `json:"is_response"`
	HideDetails                bool      `json:"hide_details"`
}

// RRModule module handler
type RRModule struct {
	DataBase *sql.DB
	RedisDB  *utils.RedisDb
}
