package models

import "time"

type PlayEvent struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	SongID    int       `json:"song_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Duration  int       `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
}

type PlayEventAggregate struct {
	SongID        int       `json:"song_id"`
	PlayCount     int       `json:"play_count"`
	TotalDuration int       `json:"total_duration"`
	LastPlayed    time.Time `json:"last_played"`
}
