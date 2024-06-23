package repositories

import (
	"errors"
	"louderspace/internal/models"
	"strings"
	"sync"
)

type StationStorageMock struct {
	stations map[int]*models.Station
	Songs    []*models.Song
	nextID   int
	mu       sync.RWMutex
}

func NewStationStorageMock() *StationStorageMock {
	return &StationStorageMock{
		stations: make(map[int]*models.Station),
		nextID:   1,
	}
}

func (t *StationStorageMock) Create(station *models.Station) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	station.ID = t.nextID
	t.nextID++
	t.stations[station.ID] = station
	return nil
}

func (t *StationStorageMock) Update(station *models.Station) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, exists := t.stations[station.ID]; !exists {
		return errors.New("station not found")
	}

	t.stations[station.ID] = station
	return nil
}

func (t *StationStorageMock) Delete(stationID int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, exists := t.stations[stationID]; !exists {
		return errors.New("station not found")
	}

	delete(t.stations, stationID)
	return nil
}

func (t *StationStorageMock) ByID(stationID int) (*models.Station, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	station, exists := t.stations[stationID]
	if !exists {
		return nil, errors.New("station not found")
	}

	return station, nil
}

func (t *StationStorageMock) All() ([]*models.Station, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var stations []*models.Station
	for _, station := range t.stations {
		stations = append(stations, station)
	}

	return stations, nil
}

func (t *StationStorageMock) SongsByTags(tags []string) ([]*models.Song, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var matchedSongs []*models.Song
	for _, song := range t.Songs {
		songTags := strings.Split(song.Genre, ",")
		for _, tag := range tags {
			if contains(songTags, tag) {
				matchedSongs = append(matchedSongs, song)
				break
			}
		}
	}

	return matchedSongs, nil
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
