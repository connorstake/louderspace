package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"louderspace/internal/services"
	"net/http"
	"strconv"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	station, err := h.stationService.CreateStation(req.Name, req.Tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
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

func (h *StationAPI) GetSongsForStationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stationID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid station ID", http.StatusBadRequest)
		return
	}
	songs, err := h.stationService.GetSongsForStation(stationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}

func (h *StationAPI) DeleteStation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stationID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid station ID", http.StatusBadRequest)
		return
	}

	err = h.stationService.DeleteStation(stationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *StationAPI) UpdateStation(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}
	stationID, err := strconv.Atoi(r.URL.Path[len("/stations/"):])
	if err != nil {
		http.Error(w, "Invalid station ID", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	station, err := h.stationService.UpdateStation(stationID, req.Name, req.Tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(station)
}
