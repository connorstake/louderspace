package repositories

import (
	"database/sql"
	"louderspace/internal/models"
	"time"
)

type PomodoroSessionStorage interface {
	Create(session *models.PomodoroSession) error
	Update(session *models.PomodoroSession) error
	ByID(id int) (*models.PomodoroSession, error)
	ByUserID(userID int) ([]*models.PomodoroSession, error)
}

type PomodoroSessionDatabase struct {
	db *sql.DB
}

func NewPomodoroSessionDatabase(db *sql.DB) PomodoroSessionStorage {
	return &PomodoroSessionDatabase{db}
}

func (r *PomodoroSessionDatabase) Create(session *models.PomodoroSession) error {
	query := "INSERT INTO pomodoro_sessions (user_id, start_time, duration, break_duration, status) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	return r.db.QueryRow(query, session.UserID, time.Now(), session.Duration, session.BreakDuration, session.Status).Scan(&session.ID)
}

func (r *PomodoroSessionDatabase) Update(session *models.PomodoroSession) error {
	query := "UPDATE pomodoro_sessions SET end_time = $1, status = $2 WHERE id = $3"
	_, err := r.db.Exec(query, time.Now(), session.Status, session.ID)
	return err
}

func (r *PomodoroSessionDatabase) ByID(id int) (*models.PomodoroSession, error) {
	session := &models.PomodoroSession{}
	query := "SELECT id, user_id, start_time, end_time, duration, break_duration, status FROM pomodoro_sessions WHERE id = $1"
	if err := r.db.QueryRow(query, id).Scan(&session.ID, &session.UserID, &session.StartTime, &session.EndTime, &session.Duration, &session.BreakDuration, &session.Status); err != nil {
		return nil, err
	}
	return session, nil
}

func (r *PomodoroSessionDatabase) ByUserID(userID int) ([]*models.PomodoroSession, error) {
	query := "SELECT id, user_id, start_time, end_time, duration, break_duration, status FROM pomodoro_sessions WHERE user_id = $1"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*models.PomodoroSession
	for rows.Next() {
		session := &models.PomodoroSession{}
		if err := rows.Scan(&session.ID, &session.UserID, &session.StartTime, &session.EndTime, &session.Duration, &session.BreakDuration, &session.Status); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}
