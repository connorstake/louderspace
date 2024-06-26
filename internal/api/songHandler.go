package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"louderspace/internal/logger"
	"louderspace/internal/models"
	"louderspace/internal/services"
	"net/http"
	"strconv"
)

type SongAPI struct {
	songService services.SongManagement
}

func NewSongAPI(songService services.SongManagement) *SongAPI {
	return &SongAPI{songService}
}

func (h *SongAPI) CreateSong(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string   `json:"title"`
		Artist      string   `json:"artist"`
		Genre       string   `json:"genre"`
		SunoID      string   `json:"suno_id"`
		IsGenerated bool     `json:"is_generated"`
		Tags        []string `json:"tags"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	song, err := h.songService.CreateSong(req.Title, req.Artist, req.Genre, req.SunoID, req.IsGenerated, req.Tags)
	if err != nil {
		logger.Error("Failed to create song:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Created song:", song)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(song)
}

func (h *SongAPI) GetSong(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		err := "Missing song ID"
		logger.Error(err)
		http.Error(w, "Missing song ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error("Invalid song ID:", err)
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	song, err := h.songService.GetSongByID(id)
	if err != nil {
		logger.Error("Failed to get song by ID:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Fetched song:", song)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(song)
}

func (h *SongAPI) GetSongBySunoID(w http.ResponseWriter, r *http.Request) {
	sunoID := r.URL.Query().Get("suno_id")
	if sunoID == "" {
		http.Error(w, "Missing Suno ID", http.StatusBadRequest)
		return
	}

	song, err := h.songService.GetSongBySunoID(sunoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(song)
}

func (h *SongAPI) GetAllSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := h.songService.GetAllSongs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}

func (h *SongAPI) UpdateSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Invalid song ID:", err)
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Title       string   `json:"title"`
		Artist      string   `json:"artist"`
		Genre       string   `json:"genre"` // Update to array
		SunoID      string   `json:"suno_id"`
		IsGenerated bool     `json:"is_generated"`
		Tags        []string `json:"tags"` // Add tags field
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	song := &models.Song{
		ID:          id,
		Title:       req.Title,
		Artist:      req.Artist,
		Genre:       req.Genre,
		SunoID:      req.SunoID,
		IsGenerated: req.IsGenerated,
	}

	if _, err := h.songService.UpdateSong(song, req.Tags); err != nil {
		logger.Error("Failed to update song:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Updated song:", song)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(song)
}

func (h *SongAPI) DeleteSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Invalid song ID:", err)
		logger.Error("Failed to delete song:", err)
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	if err := h.songService.DeleteSong(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Deleted song with ID:", id)
	w.WriteHeader(http.StatusNoContent)
}
