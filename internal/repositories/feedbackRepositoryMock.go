// mock_feedback_storage.go
package repositories

import (
	"errors"
	"louderspace/internal/models"
)

type MockFeedbackStorage struct {
	Feedbacks []models.Feedback
}

func NewMockFeedbackStorage() *MockFeedbackStorage {
	return &MockFeedbackStorage{
		Feedbacks: []models.Feedback{},
	}
}

func (m *MockFeedbackStorage) SaveFeedback(feedback *models.Feedback) error {
	for i, f := range m.Feedbacks {
		if f.UserID == feedback.UserID && f.SongID == feedback.SongID {
			m.Feedbacks[i] = *feedback
			return nil
		}
	}
	m.Feedbacks = append(m.Feedbacks, *feedback)
	return nil
}

func (m *MockFeedbackStorage) DeleteFeedback(userID, songID int) error {
	for i, f := range m.Feedbacks {
		if f.UserID == userID && f.SongID == songID {
			m.Feedbacks = append(m.Feedbacks[:i], m.Feedbacks[i+1:]...)
			return nil
		}
	}
	return errors.New("feedback not found")
}

func (m *MockFeedbackStorage) GetFeedback(userID, songID int) (*models.Feedback, error) {
	for _, f := range m.Feedbacks {
		if f.UserID == userID && f.SongID == songID {
			return &f, nil
		}
	}
	return nil, errors.New("feedback not found")
}
