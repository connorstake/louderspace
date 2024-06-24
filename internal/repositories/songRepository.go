package repositories

import (
	"database/sql"
	"louderspace/internal/models"
	"strings"
)

type SongStorage interface {
	Create(song *models.Song) error
	ByID(songID int) (*models.Song, error)
	BySunoID(sunoID string) (*models.Song, error)
	All() ([]*models.Song, error)
	SongsByTags(tags []string) ([]*models.Song, error)
	Delete(songID int) error
	Update(song *models.Song) error
}

type SongDatabase struct {
	db *sql.DB
}

func NewSongDatabase(db *sql.DB) SongStorage {
	return &SongDatabase{db}
}

func (r *SongDatabase) Create(song *models.Song) error {
	query := "INSERT INTO songs (title, artist, genre, suno_id, is_generated, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	return r.db.QueryRow(query, song.Title, song.Artist, song.Genre, song.SunoID, song.IsGenerated, song.CreatedAt).Scan(&song.ID)
}

func (r *SongDatabase) ByID(songID int) (*models.Song, error) {
	song := &models.Song{}
	query := "SELECT id, title, artist, genre, suno_id, is_generated, created_at FROM songs WHERE id=$1"
	if err := r.db.QueryRow(query, songID).Scan(&song.ID, &song.Title, &song.Artist, &song.Genre, &song.SunoID, &song.IsGenerated, &song.CreatedAt); err != nil {
		return nil, err
	}
	return song, nil
}

func (r *SongDatabase) BySunoID(sunoID string) (*models.Song, error) {
	song := &models.Song{}
	query := "SELECT id, title, artist, genre, suno_id, is_generated, created_at FROM songs WHERE suno_id=$1"
	if err := r.db.QueryRow(query, sunoID).Scan(&song.ID, &song.Title, &song.Artist, &song.Genre, &song.SunoID, &song.IsGenerated, &song.CreatedAt); err != nil {
		return nil, err
	}
	return song, nil
}

func (r *SongDatabase) All() ([]*models.Song, error) {
	var songs []*models.Song
	rows, err := r.db.Query("SELECT id, title, artist, genre, suno_id, is_generated, created_at FROM songs")
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

func (r *SongDatabase) SongsByTags(tags []string) ([]*models.Song, error) {
	var songs []*models.Song
	query := "SELECT id, title, artist, genre, suno_id, is_generated, created_at FROM songs WHERE"
	var conditions []string
	var args []interface{}
	for _, tag := range tags {
		conditions = append(conditions, "genre ILIKE ?")
		args = append(args, "%"+tag+"%")
	}
	query += strings.Join(conditions, " OR ")

	rows, err := r.db.Query(query, args...)
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

func (r *SongDatabase) Delete(songID int) error {
	_, err := r.db.Exec("DELETE FROM songs WHERE id=$1", songID)
	return err
}

func (r *SongDatabase) Update(song *models.Song) error {
	_, err := r.db.Exec("UPDATE songs SET title=$1, artist=$2, genre=$3, suno_id=$4, is_generated=$5 WHERE id=$6", song.Title, song.Artist, song.Genre, song.SunoID, song.IsGenerated, song.ID)
	return err
}
