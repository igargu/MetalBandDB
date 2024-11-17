package models

import "time"

type Band struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Genre  string    `json:"genre"`
	Formed time.Time `json:"formed"`
}
