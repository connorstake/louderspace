package repositories

import (
	"errors"
	"louderspace/internal/models"
	"sync"
)

type MockTagStorage struct {
	tags   map[int]*models.Tag
	nextID int
	mu     sync.RWMutex
}

func NewMockTagStorage() *MockTagStorage {
	return &MockTagStorage{
		tags:   make(map[int]*models.Tag),
		nextID: 1,
	}
}

func (m *MockTagStorage) GetAllTags() ([]*models.Tag, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var tags []*models.Tag
	for _, tag := range m.tags {
		tags = append(tags, tag)
	}
	return tags, nil
}

func (m *MockTagStorage) Create(tag *models.Tag) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	tag.ID = m.nextID
	m.nextID++
	m.tags[tag.ID] = tag
	return nil
}

func (m *MockTagStorage) Update(tag *models.Tag) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.tags[tag.ID]; !exists {
		return errors.New("tag not found")
	}

	m.tags[tag.ID] = tag
	return nil
}

func (m *MockTagStorage) Delete(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.tags[id]; !exists {
		return errors.New("tag not found")
	}

	delete(m.tags, id)
	return nil
}
