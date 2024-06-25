package api

import (
	"encoding/json"
	"louderspace/internal/services"
	"net/http"
	"time"
)

type PlayEventAPI struct {
	playEventService services.PlayEventManagement
}

func NewPlayEventAPI(playEventService services.PlayEventManagement) *PlayEventAPI {
	return &PlayEventAPI{playEventService}
}

func (h *PlayEventAPI) LogPlayEvent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID    int       `json:"user_id"`
		SongID    int       `json:"song_id"`
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.playEventService.LogPlayEvent(req.UserID, req.SongID, req.StartTime, req.EndTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *PlayEventAPI) UpdateAggregates(w http.ResponseWriter, r *http.Request) {
	err := h.playEventService.UpdateAggregates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
