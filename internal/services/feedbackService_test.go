// feedback_service_test.go
package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"louderspace/internal/services"
)

func TestFeedbackService_SaveFeedback(t *testing.T) {
	mockFeedbackStorage := repositories.NewMockFeedbackStorage()
	feedbackService := services.NewFeedbackService(mockFeedbackStorage)

	feedback := &models.Feedback{
		UserID: 1,
		SongID: 1,
		Liked:  true,
	}

	err := feedbackService.SaveFeedback(feedback)
	assert.NoError(t, err)

	savedFeedback, err := mockFeedbackStorage.GetFeedback(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, feedback.UserID, savedFeedback.UserID)
	assert.Equal(t, feedback.SongID, savedFeedback.SongID)
	assert.Equal(t, feedback.Liked, savedFeedback.Liked)
}

func TestFeedbackService_DeleteFeedback(t *testing.T) {
	mockFeedbackStorage := repositories.NewMockFeedbackStorage()
	feedbackService := services.NewFeedbackService(mockFeedbackStorage)

	feedback := &models.Feedback{
		UserID: 1,
		SongID: 1,
		Liked:  true,
	}

	err := mockFeedbackStorage.SaveFeedback(feedback)
	assert.NoError(t, err)

	err = feedbackService.DeleteFeedback(1, 1)
	assert.NoError(t, err)

	_, err = mockFeedbackStorage.GetFeedback(1, 1)
	assert.Error(t, err)
}

func TestFeedbackService_GetFeedback(t *testing.T) {
	mockFeedbackStorage := repositories.NewMockFeedbackStorage()
	feedbackService := services.NewFeedbackService(mockFeedbackStorage)

	feedback := &models.Feedback{
		UserID: 1,
		SongID: 1,
		Liked:  true,
	}

	err := mockFeedbackStorage.SaveFeedback(feedback)
	assert.NoError(t, err)

	retrievedFeedback, err := feedbackService.GetFeedback(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, feedback.UserID, retrievedFeedback.UserID)
	assert.Equal(t, feedback.SongID, retrievedFeedback.SongID)
	assert.Equal(t, feedback.Liked, retrievedFeedback.Liked)
}
