package services

import (
	"golang.org/x/crypto/bcrypt"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"time"
)

type UserManagement interface {
	Register(username, password, email string, role models.Role) (*models.User, error)
	Login(username, password string) (*models.User, error)
	User(userID int) (*models.User, error)
	Users() ([]*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

type UserService struct {
	userStorage repositories.UserStorage
}

func NewUserService(userStorage repositories.UserStorage) UserManagement {
	return &UserService{userStorage}
}

func (s *UserService) Register(username, password, email string, role models.Role) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:  username,
		Password:  string(hashedPassword),
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}

	if err := s.userStorage.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(username, password string) (*models.User, error) {
	user, err := s.userStorage.UserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) User(userID int) (*models.User, error) {
	return s.userStorage.UserByID(userID)
}

func (s *UserService) Users() ([]*models.User, error) {
	return s.userStorage.Users()
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.userStorage.UserByUsername(username)
}
