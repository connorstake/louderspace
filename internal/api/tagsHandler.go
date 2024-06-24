package api

import (
	"encoding/json"
	"louderspace/internal/services"
	"net/http"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}
