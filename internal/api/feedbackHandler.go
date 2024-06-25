// feedback_api.go
package api

import (
	"encoding/json"
	"louderspace/internal/logger"
	"louderspace/internal/models"
	"louderspace/internal/services"
	"net/http"
	"strconv"
)

type FeedbackAPI struct {
	feedbackService *services.FeedbackService
}

func NewFeedbackAPI(feedbackService *services.FeedbackService) *FeedbackAPI {
	return &FeedbackAPI{feedbackService}
}

func (a *FeedbackAPI) SaveFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback models.Feedback
	if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
		logger.Error("Failed to decode request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := a.feedbackService.SaveFeedback(&feedback); err != nil {
		logger.Error("Failed to save feedback:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("Saved feedback:", feedback)
	w.WriteHeader(http.StatusCreated)
}

func (a *FeedbackAPI) DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		logger.Error("Failed to parse user ID:", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	songID, err := strconv.Atoi(r.URL.Query().Get("song_id"))
	if err != nil {
		logger.Error("Failed to parse song ID:", err)
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}
	if err := a.feedbackService.DeleteFeedback(userID, songID); err != nil {
		logger.Error("Failed to delete feedback:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("Deleted feedback for user", userID, "and song", songID)
	w.WriteHeader(http.StatusNoContent)
}

func (a *FeedbackAPI) GetFeedback(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		logger.Error("Failed to parse user ID:", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	songID, err := strconv.Atoi(r.URL.Query().Get("song_id"))
	if err != nil {
		logger.Error("Failed to parse song ID:", err)
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}
	feedback, err := a.feedbackService.GetFeedback(userID, songID)
	if err != nil {
		logger.Error("Failed to get feedback:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(feedback)
	if err != nil {
		logger.Error("Failed to encode response:", err)
		return
	}
}
