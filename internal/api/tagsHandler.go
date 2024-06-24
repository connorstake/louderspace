package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"louderspace/internal/logger"
	"louderspace/internal/services"
	"net/http"
	"strconv"
)

type TagAPI struct {
	tagService services.TagManagement
}

func NewTagAPI(tagService services.TagManagement) *TagAPI {
	return &TagAPI{tagService}
}

func (h *TagAPI) GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := h.tagService.GetAllTags()
	if err != nil {
		logger.Error("Failed to get all tags:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Got all tags:", tags)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

func (h *TagAPI) CreateTag(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tag, err := h.tagService.CreateTag(req.Name)
	if err != nil {
		logger.Error("Failed to create tag:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Created tag:", tag)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tag)
}

func (h *TagAPI) UpdateTag(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Invalid tag ID:", err)
		http.Error(w, "Invalid tag ID", http.StatusBadRequest)
		return
	}

	if err := h.tagService.UpdateTag(id, req.Name); err != nil {
		logger.Error("Failed to update tag:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Updated tag with ID:", id)
	w.WriteHeader(http.StatusNoContent)
}

func (h *TagAPI) DeleteTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Invalid tag ID:", err)
		http.Error(w, "Invalid tag ID", http.StatusBadRequest)
		return
	}

	if err := h.tagService.DeleteTag(id); err != nil {
		logger.Error("Failed to delete tag:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Deleted tag with ID:", id)
	w.WriteHeader(http.StatusNoContent)
}
