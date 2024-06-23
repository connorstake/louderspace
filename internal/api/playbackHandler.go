package api

import (
	"encoding/json"
	"louderspace/internal/services"
	"net/http"
	"strconv"
)

type PlaybackAPI struct {
	playbackService services.PlaybackManagement
}

func NewPlaybackAPI(playbackService services.PlaybackManagement) *PlaybackAPI {
	return &PlaybackAPI{playbackService}
}

func (h *PlaybackAPI) Play(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	stationID, err := strconv.Atoi(r.URL.Query().Get("station_id"))
	if err != nil {
		http.Error(w, "Invalid station ID", http.StatusBadRequest)
		return
	}

	playbackState, err := h.playbackService.Play(userID, stationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playbackState)
}

func (h *PlaybackAPI) Pause(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	playbackState, err := h.playbackService.Pause(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playbackState)
}

func (h *PlaybackAPI) Skip(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	playbackState, err := h.playbackService.Skip(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playbackState)
}

func (h *PlaybackAPI) Rewind(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	playbackState, err := h.playbackService.Rewind(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playbackState)
}

func (h *PlaybackAPI) GetPlaybackState(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	playbackState, err := h.playbackService.GetPlaybackState(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playbackState)
}
