package api

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"louderspace/internal/logger"
	"louderspace/internal/middleware"
	"louderspace/internal/models"
	"louderspace/internal/services"
	"louderspace/internal/utils"
	"net/http"
	"time"
)

type AuthAPI struct {
	userService services.UserManagement
}

func NewAuthAPI(userService services.UserManagement) *AuthAPI {
	return &AuthAPI{userService}
}

func (a *AuthAPI) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string      `json:"username"`
		Password string      `json:"password"`
		Email    string      `json:"email"`
		Role     models.Role `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to hash password", err)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user, err := a.userService.Register(req.Username, string(hashedPassword), req.Email, req.Role)
	if err != nil {
		logger.Error("Failed to register user", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("User registered", user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (a *AuthAPI) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logger.Info(req.Username, req.Password)

	user, err := a.userService.GetUserByUsername(req.Username)
	if err != nil {
		logger.Error("Failed to get user", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		logger.Error("Failed to compare password", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		logger.Error("Failed to generate token", err)
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	})

	logger.Info("User logged in", user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user":  user,
	})
}

func (a *AuthAPI) Me(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.UserContextKey).(*models.User)
	if !ok {
		logger.Error("No user in context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fullUser, err := a.userService.User(user.ID)
	if err != nil {
		logger.Error("Failed to get user", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fullUser)
}
