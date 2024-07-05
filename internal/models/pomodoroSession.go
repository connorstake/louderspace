package models

import "time"

type PomodoroSession struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	Duration      int       `json:"duration"`
	BreakDuration int       `json:"break_duration"`
	Status        string    `json:"status"`
}
