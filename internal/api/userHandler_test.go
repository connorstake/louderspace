package api

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"louderspace/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	userStorage := repositories.NewMockUserStorage()
	userService := services.NewUserService(userStorage)
	userAPI := NewUserAPI(userService)

	payload := map[string]string{
		"username": "testuser",
		"password": "password123",
		"email":    "test@example.com",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userAPI.Register)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var user models.User
	err := json.Unmarshal(rr.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestLogin(t *testing.T) {
	userStorage := repositories.NewMockUserStorage()
	userService := services.NewUserService(userStorage)
	userAPI := NewUserAPI(userService)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	userStorage.Save(&models.User{
		Username:  "testuser",
		Password:  string(hashedPassword),
		Email:     "test@example.com",
		CreatedAt: time.Now(),
	})

	payload := map[string]string{
		"username": "testuser",
		"password": "password123",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userAPI.Login)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var user models.User
	err := json.Unmarshal(rr.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
}
