package services

import (
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"time"
)

type PomodoroSessionManagement interface {
	StartSession(session *models.PomodoroSession) error
	EndSession(sessionID int) error
	GetSessionsByUser(userID int) ([]*models.PomodoroSession, error)
	GetFocusMetrics(userID int) (int, int, error)
}

type PomodoroSessionService struct {
	pomodoroRepo repositories.PomodoroSessionStorage
}

func NewPomodoroSessionService(pomodoroRepo repositories.PomodoroSessionStorage) PomodoroSessionManagement {
	return &PomodoroSessionService{pomodoroRepo}
}

func (s *PomodoroSessionService) StartSession(session *models.PomodoroSession) error {
	return s.pomodoroRepo.Create(session)
}

func (s *PomodoroSessionService) EndSession(sessionID int) error {
	session, err := s.pomodoroRepo.ByID(sessionID)
	if err != nil {
		return err
	}
	session.EndTime = time.Now()
	session.Status = "completed"
	return s.pomodoroRepo.Update(session)
}

func (s *PomodoroSessionService) GetSessionsByUser(userID int) ([]*models.PomodoroSession, error) {
	return s.pomodoroRepo.ByUserID(userID)
}

func (s *PomodoroSessionService) GetFocusMetrics(userID int) (int, int, error) {
	sessions, err := s.pomodoroRepo.ByUserID(userID)
	if err != nil {
		return 0, 0, err
	}

	var totalFocusTime int
	var sessionCount int
	for _, session := range sessions {
		if session.EndTime.After(session.StartTime) {
			focusDuration := session.EndTime.Sub(session.StartTime)
			totalFocusTime += int(focusDuration.Minutes())
			sessionCount++
		}
	}
	return totalFocusTime, sessionCount, nil
}
