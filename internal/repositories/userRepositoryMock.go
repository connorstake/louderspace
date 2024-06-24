package repositories

import (
	"errors"
	"louderspace/internal/models"
	"sync"
)

type MockUserStorage struct {
	users  map[int]*models.User
	mu     sync.RWMutex
	nextID int
}

func NewMockUserStorage() *MockUserStorage {
	return &MockUserStorage{
		users:  make(map[int]*models.User),
		nextID: 1,
	}
}

func (s *MockUserStorage) Save(user *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if user.ID == 0 {
		user.ID = s.nextID
		s.nextID++
	} else if _, exists := s.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	s.users[user.ID] = user
	return nil
}

func (s *MockUserStorage) UserByID(userID int) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[userID]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *MockUserStorage) UserByUsername(username string) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (s *MockUserStorage) Users() ([]*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var users []*models.User
	for _, user := range s.users {
		users = append(users, user)
	}

	return users, nil
}
