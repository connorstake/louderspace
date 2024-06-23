package api

import (
	"encoding/json"
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
		Title       string `json:"title"`
		Artist      string `json:"artist"`
		Genre       string `json:"genre"`
		SunoID      string `json:"suno_id"`
		IsGenerated bool   `json:"is_generated"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	song, err := h.songService.CreateSong(req.Title, req.Artist, req.Genre, req.SunoID, req.IsGenerated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(song)
}

func (h *SongAPI) GetSong(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing song ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	song, err := h.songService.GetSongByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
