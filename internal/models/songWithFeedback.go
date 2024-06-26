package models

type SongWithFeedback struct {
	*Song
	Liked bool `json:"liked"`
}
