package services

import (
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"time"
)

type PlayEventManagement interface {
	LogPlayEvent(userID, songID int, startTime, endTime time.Time) error
	UpdateAggregates() error
}

type PlayEventService struct {
	playEventStorage repositories.PlayEventStorage
}

func NewPlayEventService(playEventStorage repositories.PlayEventStorage) PlayEventManagement {
	return &PlayEventService{playEventStorage}
}

func (s *PlayEventService) LogPlayEvent(userID, songID int, startTime, endTime time.Time) error {
	event := &models.PlayEvent{
		UserID:    userID,
		SongID:    songID,
		StartTime: startTime,
		EndTime:   endTime,
		CreatedAt: time.Now(),
	}
	return s.playEventStorage.LogPlayEvent(event)
}

func (s *PlayEventService) UpdateAggregates() error {
	return s.playEventStorage.UpdateAggregates()
}
