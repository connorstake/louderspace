package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"louderspace/internal/logger"
	"louderspace/internal/models"
	"louderspace/internal/services"
	"net/http"
	"strconv"
	"time"
)

type PomodoroSessionAPI struct {
	pomodoroService services.PomodoroSessionManagement
}

func NewPomodoroSessionAPI(pomodoroService services.PomodoroSessionManagement) *PomodoroSessionAPI {
	return &PomodoroSessionAPI{pomodoroService}
}

func (h *PomodoroSessionAPI) StartSession(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID int `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session := &models.PomodoroSession{
		UserID:    req.UserID,
		StartTime: time.Now(),
		Status:    "ongoing",
	}

	if err := h.pomodoroService.StartSession(session); err != nil {
		logger.Error("Failed to start pomodoro session:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Started pomodoro session:", session)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(session)
	if err != nil {
		logger.Error("Failed to encode response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PomodoroSessionAPI) EndSession(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SessionID int `json:"session_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.pomodoroService.EndSession(req.SessionID); err != nil {
		logger.Error("Failed to end pomodoro session:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Ended pomodoro session:", req.SessionID)
	w.WriteHeader(http.StatusNoContent)
}

func (h *PomodoroSessionAPI) GetSessionsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		logger.Error("Invalid user ID:", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	sessions, err := h.pomodoroService.GetSessionsByUser(userID)
	if err != nil {
		logger.Error("Failed to get pomodoro sessions for user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Got pomodoro sessions for user:", sessions)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(sessions)
	if err != nil {
		logger.Error("Failed to encode response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PomodoroSessionAPI) GetFocusMetrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		logger.Error("Invalid user ID:", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	totalFocusTime, sessionCount, err := h.pomodoroService.GetFocusMetrics(userID)
	if err != nil {
		logger.Error("Failed to get focus metrics for user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]int{
		"total_focus_time": totalFocusTime,
		"session_count":    sessionCount,
	}

	logger.Info("Got focus metrics for user:", response)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Error("Failed to encode response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
