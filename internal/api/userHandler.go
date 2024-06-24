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

func (h *UserAPI) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.Register(req.Username, req.Password, req.Email)
	if err != nil {
		logger.Error("Failed to register user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Registered user:", user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserAPI) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.Login(req.Username, req.Password)
	if err != nil {
		logger.Error("Failed to login user:", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	logger.Info("Logged in user:", user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
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
