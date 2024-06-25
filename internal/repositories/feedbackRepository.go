package repositories

import (
	"database/sql"
	"louderspace/internal/models"
)

type FeedbackStorage interface {
	SaveFeedback(feedback *models.Feedback) error
	DeleteFeedback(userID, songID int) error
	GetFeedback(userID, songID int) (*models.Feedback, error)
}

type FeedbackDatabase struct {
	db *sql.DB
}

func NewFeedbackDatabase(db *sql.DB) FeedbackStorage {
	return &FeedbackDatabase{db}
}

func (f *FeedbackDatabase) SaveFeedback(feedback *models.Feedback) error {
	query := "INSERT INTO feedback (user_id, song_id, liked) VALUES ($1, $2, $3) ON CONFLICT (user_id, song_id) DO UPDATE SET liked = $3"
	_, err := f.db.Exec(query, feedback.UserID, feedback.SongID, feedback.Liked)
	return err
}

func (f *FeedbackDatabase) DeleteFeedback(userID, songID int) error {
	query := "DELETE FROM feedback WHERE user_id = $1 AND song_id = $2"
	_, err := f.db.Exec(query, userID, songID)
	return err
}

func (f *FeedbackDatabase) GetFeedback(userID, songID int) (*models.Feedback, error) {
	feedback := &models.Feedback{}
	query := "SELECT id, user_id, song_id, liked FROM feedback WHERE user_id = $1 AND song_id = $2"
	err := f.db.QueryRow(query, userID, songID).Scan(&feedback.ID, &feedback.UserID, &feedback.SongID, &feedback.Liked)
	if err != nil {
		return nil, err
	}
	return feedback, nil
}
