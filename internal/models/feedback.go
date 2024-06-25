package models

type Feedback struct {
	ID     int  `json:"id"`
	UserID int  `json:"user_id"`
	SongID int  `json:"song_id"`
	Liked  bool `json:"liked"`
}
