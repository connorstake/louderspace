package repositories

import (
	"database/sql"
	"louderspace/internal/models"
)

type PlayEventStorage interface {
	LogPlayEvent(event *models.PlayEvent) error
	UpdateAggregates() error
}

type PlayEventDatabase struct {
	db *sql.DB
}

func NewPlayEventDatabase(db *sql.DB) PlayEventStorage {
	return &PlayEventDatabase{db}
}

func (r *PlayEventDatabase) LogPlayEvent(event *models.PlayEvent) error {
	query := "INSERT INTO play_events (user_id, song_id, start_time, end_time, duration, created_at) VALUES ($1, $2, $3, $4, EXTRACT(EPOCH FROM $4 - $3), $5)"
	_, err := r.db.Exec(query, event.UserID, event.SongID, event.StartTime, event.EndTime, event.CreatedAt)
	return err
}

func (r *PlayEventDatabase) UpdateAggregates() error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	updateQuery := `
        INSERT INTO song_play_counts (song_id, play_count, total_duration, last_played)
        SELECT song_id, COUNT(*), SUM(duration), MAX(end_time)
        FROM play_events
        GROUP BY song_id
        ON CONFLICT (song_id) DO UPDATE
        SET play_count = EXCLUDED.play_count,
            total_duration = EXCLUDED.total_duration,
            last_played = EXCLUDED.last_played
    `
	_, err = tx.Exec(updateQuery)
	return err
}
