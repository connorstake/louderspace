package api

import (
	"encoding/json"
	"louderspace/internal/logger"
	"louderspace/internal/services"
	"net/http"
)

type UserAPI struct {
	userService services.UserManagement
}

func NewUserAPI(userService services.UserManagement) *UserAPI {
	return &UserAPI{userService}
}

func (h *UserAPI) Users(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.Users()
	if err != nil {
		logger.Error("Failed to get all users:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Got all users:", users)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
