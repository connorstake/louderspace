package models

import "time"

type PlaybackState struct {
	UserID       int       `json:"user_id"`
	StationID    int       `json:"station_id"`
	CurrentSong  *Song     `json:"current_song"`
	SongQueue    []*Song   `json:"song_queue"`
	PlaybackTime time.Time `json:"playback_time"`
	IsPlaying    bool      `json:"is_playing"`
}
