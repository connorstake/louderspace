package repositories

import (
	"database/sql"
	"louderspace/internal/models"
)

type SongStorage interface {
	Create(song *models.Song, tags []string) error
	Update(song *models.Song, tags []string) error
	ByID(id int) (*models.Song, error)
	BySunoID(sunoID string) (*models.Song, error)
	All() ([]*models.Song, error)
	ByStationID(stationID int) ([]*models.Song, error)
	Delete(id int) error
}

type SongDatabase struct {
	db *sql.DB
}

func NewSongDatabase(db *sql.DB) SongStorage {
	return &SongDatabase{db}
}

func (r *SongDatabase) Create(song *models.Song, tags []string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO songs (title, artist, genre, suno_id, is_generated, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err = tx.QueryRow(query, song.Title, song.Artist, song.Genre, song.SunoID, song.IsGenerated, song.CreatedAt).Scan(&song.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, tag := range tags {
		_, err := tx.Exec("INSERT INTO song_tags (song_id, tag_id) VALUES ($1, (SELECT id FROM tags WHERE name = $2))", song.ID, tag)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *SongDatabase) Update(song *models.Song, tags []string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := "UPDATE songs SET title = $1, artist = $2, genre = $3, suno_id = $4, is_generated = $5 WHERE id = $6"
	_, err = tx.Exec(query, song.Title, song.Artist, song.Genre, song.SunoID, song.IsGenerated, song.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM song_tags WHERE song_id = $1", song.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, tag := range tags {
		_, err := tx.Exec("INSERT INTO song_tags (song_id, tag_id) VALUES ($1, (SELECT id FROM tags WHERE name = $2))", song.ID, tag)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *SongDatabase) ByID(id int) (*models.Song, error) {
	song := &models.Song{}
	query := "SELECT id, title, artist, genre, suno_id, is_generated, created_at FROM songs WHERE id = $1"
	if err := r.db.QueryRow(query, id).Scan(&song.ID, &song.Title, &song.Artist, &song.Genre, &song.SunoID, &song.IsGenerated, &song.CreatedAt); err != nil {
		return nil, err
	}
	return song, nil
}

func (r *SongDatabase) BySunoID(sunoID string) (*models.Song, error) {
	song := &models.Song{}
	query := "SELECT id, title, artist, genre, suno_id, is_generated, created_at FROM songs WHERE suno_id = $1"
	if err := r.db.QueryRow(query, sunoID).Scan(&song.ID, &song.Title, &song.Artist, &song.Genre, &song.SunoID, &song.IsGenerated, &song.CreatedAt); err != nil {
		return nil, err
	}
	return song, nil
}

func (r *SongDatabase) All() ([]*models.Song, error) {
	var songs []*models.Song
	query := "SELECT id, title, artist, genre, suno_id, is_generated, created_at FROM songs"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		song := &models.Song{}
		if err := rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Genre, &song.SunoID, &song.IsGenerated, &song.CreatedAt); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (r *SongDatabase) ByStationID(stationID int) ([]*models.Song, error) {
	var songs []*models.Song
	query := `SELECT s.id, s.title, s.artist, s.genre, s.suno_id, s.is_generated, s.created_at 
			  FROM songs s 
			  JOIN song_tags st ON s.id = st.song_id 
			  JOIN stations_tags stt ON st.tag_id = stt.tag_id 
			  WHERE stt.station_id = $1`
	rows, err := r.db.Query(query, stationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		song := &models.Song{}
		if err := rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Genre, &song.SunoID, &song.IsGenerated, &song.CreatedAt); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (r *SongDatabase) Delete(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Clean up related data in the song_tags table if necessary
	_, err = tx.Exec("DELETE FROM song_tags WHERE song_id = $1", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	query := "DELETE FROM songs WHERE id = $1"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
