package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"louderspace/internal/logger"
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
		logger.Error("Failed to decode request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	station, err := h.stationService.CreateStation(req.Name, req.Tags)
	if err != nil {
		logger.Error("Failed to create station:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Created station:", station)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(station)
	if err != nil {
		logger.Error("Failed to encode response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *StationAPI) GetAllStations(w http.ResponseWriter, _ *http.Request) {

	stations, err := h.stationService.GetAllStations()
	if err != nil {
		logger.Error("Failed to get all stations:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("Got all stations:", stations)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(stations)
	if err != nil {
		logger.Error("Failed to encode response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *StationAPI) GetSongsForStationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stationID, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Invalid station ID:", err)
		http.Error(w, "Invalid station ID", http.StatusBadRequest)
		return
	}
	songs, err := h.stationService.GetSongsForStation(stationID)
	if err != nil {
		logger.Error("Failed to get songs for station:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("Got songs for station:", songs)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		logger.Error("Failed to encode response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *StationAPI) DeleteStation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stationID, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Invalid station ID:", err)
		http.Error(w, "Invalid station ID", http.StatusBadRequest)
		return
	}

	err = h.stationService.DeleteStation(stationID)
	if err != nil {
		logger.Error("Failed to delete station:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Deleted station:", stationID)
	w.WriteHeader(http.StatusNoContent)
}

func (h *StationAPI) UpdateStation(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}
	stationID, err := strconv.Atoi(r.URL.Path[len("/admin/stations/"):])
	if err != nil {
		logger.Error("Invalid station ID:", err)
		http.Error(w, "Invalid station ID", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	station, err := h.stationService.UpdateStation(stationID, req.Name, req.Tags)
	if err != nil {
		logger.Error("Failed to update station:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Updated station:", station)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(station)
	if err != nil {
		logger.Error("Failed to encode response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
