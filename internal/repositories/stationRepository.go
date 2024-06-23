package repositories

import (
	"database/sql"
	"louderspace/internal/models"
	"strings"
)

type StationStorage interface {
	Create(station *models.Station) error
	Update(station *models.Station) error
	Delete(stationID int) error
	ByID(stationID int) (*models.Station, error)
	All() ([]*models.Station, error)
	SongsByTags(tags []string) ([]*models.Song, error)
}

type StationDatabase struct {
	db *sql.DB
}

func NewStationDatabase(db *sql.DB) StationStorage {
	return &StationDatabase{db}
}

func (r *StationDatabase) Create(station *models.Station) error {
	query := "INSERT INTO stations (name, tags) VALUES ($1, $2) RETURNING id"
	return r.db.QueryRow(query, station.Name, strings.Join(station.Tags, ",")).Scan(&station.ID)
}

func (r *StationDatabase) Update(station *models.Station) error {
	query := "UPDATE stations SET name=$1, tags=$2 WHERE id=$3"
	_, err := r.db.Exec(query, station.Name, strings.Join(station.Tags, ","), station.ID)
	return err
}

func (r *StationDatabase) Delete(stationID int) error {
	query := "DELETE FROM stations WHERE id=$1"
	_, err := r.db.Exec(query, stationID)
	return err
}

func (r *StationDatabase) ByID(stationID int) (*models.Station, error) {
	station := &models.Station{}
	var tags string
	query := "SELECT id, name, tags FROM stations WHERE id=$1"
	if err := r.db.QueryRow(query, stationID).Scan(&station.ID, &station.Name, &tags); err != nil {
		return nil, err
	}
	station.Tags = strings.Split(tags, ",")
	return station, nil
}

func (r *StationDatabase) All() ([]*models.Station, error) {
	var stations []*models.Station
	rows, err := r.db.Query("SELECT id, name, tags FROM stations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		station := &models.Station{}
		var tags string
		if err := rows.Scan(&station.ID, &station.Name, &tags); err != nil {
			return nil, err
		}
		station.Tags = strings.Split(tags, ",")
		stations = append(stations, station)
	}

	return stations, nil
}

func (r *StationDatabase) SongsByTags(tags []string) ([]*models.Song, error) {
	var songs []*models.Song
	query := "SELECT id, title, artist, genre, suno_api_id, is_generated, created_at FROM songs WHERE"
	var conditions []string
	var args []interface{}
	for _, tag := range tags {
		conditions = append(conditions, "tags ILIKE ?")
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
		if err := rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Genre, &song.SunoAPIID, &song.IsGenerated, &song.CreatedAt); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}
