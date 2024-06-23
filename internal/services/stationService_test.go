package services

import (
	"github.com/stretchr/testify/assert"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"testing"
)

func TestCreateStation(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewStationService(storage)

	station, err := service.CreateStation("Chill Beats", []string{"chill", "beats"})
	assert.NoError(t, err)
	assert.NotNil(t, station)
	assert.Equal(t, "Chill Beats", station.Name)
	assert.ElementsMatch(t, []string{"chill", "beats"}, station.Tags)
}

func TestUpdateStation(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewStationService(storage)

	station, err := service.CreateStation("Chill Beats", []string{"chill", "beats"})
	assert.NoError(t, err)

	updatedStation, err := service.UpdateStation(station.ID, "Chill Vibes", []string{"chill", "vibes"})
	assert.NoError(t, err)
	assert.NotNil(t, updatedStation)
	assert.Equal(t, "Chill Vibes", updatedStation.Name)
	assert.ElementsMatch(t, []string{"chill", "vibes"}, updatedStation.Tags)
}

func TestDeleteStation(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewStationService(storage)

	station, err := service.CreateStation("Chill Beats", []string{"chill", "beats"})
	assert.NoError(t, err)

	err = service.DeleteStation(station.ID)
	assert.NoError(t, err)

	_, err = service.GetStation(station.ID)
	assert.Error(t, err)
}

func TestGetStation(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewStationService(storage)

	station, err := service.CreateStation("Chill Beats", []string{"chill", "beats"})
	assert.NoError(t, err)

	fetchedStation, err := service.GetStation(station.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedStation)
	assert.Equal(t, "Chill Beats", fetchedStation.Name)
	assert.ElementsMatch(t, []string{"chill", "beats"}, fetchedStation.Tags)
}

func TestGetAllStations(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewStationService(storage)

	_, err := service.CreateStation("Chill Beats", []string{"chill", "beats"})
	assert.NoError(t, err)
	_, err = service.CreateStation("Lo-fi Hip Hop", []string{"lo-fi", "hip hop"})
	assert.NoError(t, err)

	stations, err := service.GetAllStations()
	assert.NoError(t, err)
	assert.Len(t, stations, 2)
}

func TestGetSongsForStation(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewStationService(storage)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
		{ID: 2, Title: "Chill Song 2", Artist: "Artist 2", Genre: "chill, vibes"},
		{ID: 3, Title: "Lo-fi Song 1", Artist: "Artist 3", Genre: "lo-fi, hip hop"},
	}

	station, err := service.CreateStation("Chill Beats", []string{"chill", "beats"})
	assert.NoError(t, err)

	songs, err := service.GetSongsForStation(station.ID)
	assert.NoError(t, err)
	assert.Len(t, songs, 2)
	assert.ElementsMatch(t, []string{"Chill Song 1", "Chill Song 2"}, []string{songs[0].Title, songs[1].Title})
}
