package api

import (
	"encoding/json"
	"log"
	"louderspace/internal/services"
	"net/http"
	"strconv"
	"strings"
)

type StationAPI struct {
	stationService services.StationManagement
}

func NewStationAPI(stationService services.StationManagement) *StationAPI {
	return &StationAPI{stationService}
}

func (h *StationAPI) CreateStation(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("CreateStation request: %v", req)

	station, err := h.stationService.CreateStation(req.Name, req.Tags)
	if err != nil {
		log.Printf("Error creating station: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(station); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *StationAPI) UpdateStation(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/stations/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid station ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	station, err := h.stationService.UpdateStation(id, req.Name, req.Tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(station)
}

func (h *StationAPI) DeleteStation(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/stations/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid station ID", http.StatusBadRequest)
		return
	}

	if err := h.stationService.DeleteStation(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *StationAPI) GetStation(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/stations/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid station ID", http.StatusBadRequest)
		return
	}

	station, err := h.stationService.GetStation(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(station)
}

func (h *StationAPI) GetAllStations(w http.ResponseWriter, r *http.Request) {
	stations, err := h.stationService.GetAllStations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stations)
}

func (h *StationAPI) GetSongsForStationByID(w http.ResponseWriter, r *http.Request, stationID int) {
	songs, err := h.stationService.GetSongsForStation(stationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}
