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
	GetTagsBySongID(songID int) ([]models.Tag, error)
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
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := "UPDATE songs SET title = $1, artist = $2, genre = $3, suno_id = $4, is_generated = $5 WHERE id = $6"
	_, err = tx.Exec(query, song.Title, song.Artist, song.Genre, song.SunoID, song.IsGenerated, song.ID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM song_tags WHERE song_id = $1", song.ID)
	if err != nil {
		return err
	}

	for _, tag := range tags {
		_, err := tx.Exec("INSERT INTO song_tags (song_id, tag_id) VALUES ($1, (SELECT id FROM tags WHERE name = $2))", song.ID, tag)
		if err != nil {
			return err
		}
	}

	return nil
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
	query := `
	SELECT s.id, s.title, s.artist, s.genre, s.suno_id, s.is_generated, s.created_at, t.id, t.name
	FROM songs s
	LEFT JOIN song_tags st ON s.id = st.song_id
	LEFT JOIN tags t ON st.tag_id = t.id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	songMap := make(map[int]*models.Song)
	tagMap := make(map[int]map[int]*models.Tag)

	for rows.Next() {
		var songID int
		var tagID sql.NullInt64
		var song models.Song
		var tag models.Tag

		err := rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Genre, &song.SunoID, &song.IsGenerated, &song.CreatedAt, &tagID, &tag.Name)
		if err != nil {
			return nil, err
		}

		songID = song.ID
		if _, exists := songMap[songID]; !exists {
			songMap[songID] = &song
			tagMap[songID] = make(map[int]*models.Tag)
		}

		if tagID.Valid {
			tag.ID = int(tagID.Int64)
			tagMap[songID][tag.ID] = &tag
		}
	}

	var songs []*models.Song
	for id, song := range songMap {
		for _, tag := range tagMap[id] {
			song.Tags = append(song.Tags, *tag)
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (r *SongDatabase) ByStationID(stationID int) ([]*models.Song, error) {
	var songs []*models.Song
	query := `
		SELECT DISTINCT s.id, s.title, s.artist, s.genre, s.suno_id, s.is_generated, s.created_at
		FROM songs s
		JOIN song_tags st ON s.id = st.song_id
		JOIN tags t ON st.tag_id = t.id
		JOIN stations stn ON stn.tags LIKE '%' || t.name || '%'
		WHERE stn.id = $1
	`
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

		// Fetch tags for the song
		tagsQuery := `
			SELECT t.id, t.name
			FROM tags t
			JOIN song_tags st ON t.id = st.tag_id
			WHERE st.song_id = $1
		`
		tagRows, err := r.db.Query(tagsQuery, song.ID)
		if err != nil {
			return nil, err
		}
		defer tagRows.Close()

		for tagRows.Next() {
			tag := models.Tag{}
			if err := tagRows.Scan(&tag.ID, &tag.Name); err != nil {
				return nil, err
			}
			song.Tags = append(song.Tags, tag)
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

func (r *SongDatabase) GetTagsBySongID(songID int) ([]models.Tag, error) {
	var tags []models.Tag
	query := `
		SELECT t.id, t.name
		FROM tags t
		JOIN song_tags st ON t.id = st.tag_id
		WHERE st.song_id = $1
	`
	rows, err := r.db.Query(query, songID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
