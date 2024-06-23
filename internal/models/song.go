package models

import "time"

type Song struct {
	ID          int       `json:"id"`
	SunoID      string    `json:"suno_id"`
	Title       string    `json:"title"`
	Artist      string    `json:"artist"`
	Genre       string    `json:"genre"`
	IsGenerated bool      `json:"is_generated"`
	CreatedAt   time.Time `json:"created_at"`
	Tags        []Tag     `json:"tags"`
}
