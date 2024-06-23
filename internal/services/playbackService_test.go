package services

import (
	"github.com/stretchr/testify/assert"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"testing"
)

func TestPlaybackService_Play(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewPlaybackService(storage)

	station := &models.Station{
		Name: "Chill Beats",
		Tags: []string{"chill", "beats"},
	}
	storage.Create(station)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
		{ID: 2, Title: "Chill Song 2", Artist: "Artist 2", Genre: "chill, vibes"},
	}

	playbackState, err := service.Play(1, station.ID)
	assert.NoError(t, err)
	assert.NotNil(t, playbackState)
	assert.True(t, playbackState.IsPlaying)
	assert.Equal(t, "Chill Song 1", playbackState.CurrentSong.Title)
}

func TestPlaybackService_Pause(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewPlaybackService(storage)

	station := &models.Station{
		Name: "Chill Beats",
		Tags: []string{"chill", "beats"},
	}
	storage.Create(station)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
		{ID: 2, Title: "Chill Song 2", Artist: "Artist 2", Genre: "chill, vibes"},
	}

	service.Play(1, station.ID)
	playbackState, err := service.Pause(1)
	assert.NoError(t, err)
	assert.NotNil(t, playbackState)
	assert.False(t, playbackState.IsPlaying)
}

func TestPlaybackService_Skip(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewPlaybackService(storage)

	station := &models.Station{
		Name: "Chill Beats",
		Tags: []string{"chill", "beats"},
	}
	storage.Create(station)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
		{ID: 2, Title: "Chill Song 2", Artist: "Artist 2", Genre: "chill, vibes"},
	}

	service.Play(1, station.ID)
	playbackState, err := service.Skip(1)
	assert.NoError(t, err)
	assert.NotNil(t, playbackState)
	assert.Equal(t, "Chill Song 2", playbackState.CurrentSong.Title)
}

func TestPlaybackService_Rewind(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewPlaybackService(storage)

	station := &models.Station{
		Name: "Chill Beats",
		Tags: []string{"chill", "beats"},
	}
	storage.Create(station)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
	}

	service.Play(1, station.ID)
	playbackState, err := service.Rewind(1)
	assert.NoError(t, err)
	assert.NotNil(t, playbackState)
	assert.Equal(t, "Chill Song 1", playbackState.CurrentSong.Title)
}

func TestPlaybackService_GetPlaybackState(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	service := NewPlaybackService(storage)

	station := &models.Station{
		Name: "Chill Beats",
		Tags: []string{"chill", "beats"},
	}
	storage.Create(station)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
	}

	service.Play(1, station.ID)
	playbackState, err := service.GetPlaybackState(1)
	assert.NoError(t, err)
	assert.NotNil(t, playbackState)
	assert.Equal(t, "Chill Song 1", playbackState.CurrentSong.Title)
}
