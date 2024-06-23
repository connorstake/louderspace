package services

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	userStorage := repositories.NewMockUserStorage()
	userService := NewUserService(userStorage)

	user, err := userService.Register("testuser", "password123", "test@example.com")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
	assert.NotEqual(t, "password123", user.Password)
}

func TestLogin(t *testing.T) {
	userStorage := repositories.NewMockUserStorage()
	userService := NewUserService(userStorage)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	userStorage.Save(&models.User{
		Username:  "testuser",
		Password:  string(hashedPassword),
		Email:     "test@example.com",
		CreatedAt: time.Now(),
	})

	user, err := userService.Login("testuser", "password123")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestUser(t *testing.T) {
	userStorage := repositories.NewMockUserStorage()
	userService := NewUserService(userStorage)

	savedUser := &models.User{
		Username:  "testuser",
		Password:  "password123",
		Email:     "test@example.com",
		CreatedAt: time.Now(),
	}
	userStorage.Save(savedUser)

	user, err := userService.User(savedUser.ID)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
}
