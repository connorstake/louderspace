// feedback_service.go
package services

import (
	"louderspace/internal/models"
	"louderspace/internal/repositories"
)

type FeedbackService struct {
	storage repositories.FeedbackStorage
}

func NewFeedbackService(storage repositories.FeedbackStorage) *FeedbackService {
	return &FeedbackService{storage}
}

func (s *FeedbackService) SaveFeedback(feedback *models.Feedback) error {
	return s.storage.SaveFeedback(feedback)
}

func (s *FeedbackService) DeleteFeedback(userID, songID int) error {
	return s.storage.DeleteFeedback(userID, songID)
}

func (s *FeedbackService) GetFeedback(userID, songID int) (*models.Feedback, error) {
	return s.storage.GetFeedback(userID, songID)
}
