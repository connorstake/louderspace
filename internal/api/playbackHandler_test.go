package api

import (
	"encoding/json"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"louderspace/internal/services"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaybackAPI_Play(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	playbackService := services.NewPlaybackService(storage)
	playbackAPI := NewPlaybackAPI(playbackService)

	station := &models.Station{
		Name: "Chill Beats",
		Tags: []string{"chill", "beats"},
	}
	storage.Create(station)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
		{ID: 2, Title: "Chill Song 2", Artist: "Artist 2", Genre: "chill, vibes"},
	}

	req, err := http.NewRequest("GET", "/playback/play?user_id=1&station_id="+strconv.Itoa(station.ID), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	http.HandlerFunc(playbackAPI.Play).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var playbackState models.PlaybackState
	err = json.NewDecoder(rr.Body).Decode(&playbackState)
	assert.NoError(t, err)
	assert.True(t, playbackState.IsPlaying)
	assert.Equal(t, "Chill Song 1", playbackState.CurrentSong.Title)
}

func TestPlaybackAPI_Pause(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	playbackService := services.NewPlaybackService(storage)
	playbackAPI := NewPlaybackAPI(playbackService)

	station := &models.Station{
		Name: "Chill Beats",
		Tags: []string{"chill", "beats"},
	}
	storage.Create(station)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
		{ID: 2, Title: "Chill Song 2", Artist: "Artist 2", Genre: "chill, vibes"},
	}

	playbackService.Play(1, station.ID)
	req, err := http.NewRequest("GET", "/playback/pause?user_id=1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	http.HandlerFunc(playbackAPI.Pause).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var playbackState models.PlaybackState
	err = json.NewDecoder(rr.Body).Decode(&playbackState)
	assert.NoError(t, err)
	assert.False(t, playbackState.IsPlaying)
}

func TestPlaybackAPI_Skip(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	playbackService := services.NewPlaybackService(storage)
	playbackAPI := NewPlaybackAPI(playbackService)

	station := &models.Station{
		Name: "Chill Beats",
		Tags: []string{"chill", "beats"},
	}
	storage.Create(station)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
		{ID: 2, Title: "Chill Song 2", Artist: "Artist 2", Genre: "chill, vibes"},
	}

	playbackService.Play(1, station.ID)
	req, err := http.NewRequest("GET", "/playback/skip?user_id=1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	http.HandlerFunc(playbackAPI.Skip).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var playbackState models.PlaybackState
	err = json.NewDecoder(rr.Body).Decode(&playbackState)
	assert.NoError(t, err)
	assert.Equal(t, "Chill Song 2", playbackState.CurrentSong.Title)
}

func TestPlaybackAPI_Rewind(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	playbackService := services.NewPlaybackService(storage)
	playbackAPI := NewPlaybackAPI(playbackService)

	station := &models.Station{
		Name: "Chill Beats",
		Tags: []string{"chill", "beats"},
	}
	storage.Create(station)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
	}

	playbackService.Play(1, station.ID)
	req, err := http.NewRequest("GET", "/playback/rewind?user_id=1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	http.HandlerFunc(playbackAPI.Rewind).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var playbackState models.PlaybackState
	err = json.NewDecoder(rr.Body).Decode(&playbackState)
	assert.NoError(t, err)
	assert.Equal(t, "Chill Song 1", playbackState.CurrentSong.Title)
}

func TestPlaybackAPI_GetPlaybackState(t *testing.T) {
	storage := repositories.NewStationStorageMock()
	playbackService := services.NewPlaybackService(storage)
	playbackAPI := NewPlaybackAPI(playbackService)

	station := &models.Station{
		Name: "Chill Beats",
		Tags: []string{"chill", "beats"},
	}
	storage.Create(station)

	storage.Songs = []*models.Song{
		{ID: 1, Title: "Chill Song 1", Artist: "Artist 1", Genre: "chill, beats"},
	}

	playbackService.Play(1, station.ID)
	req, err := http.NewRequest("GET", "/playback/state?user_id=1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	http.HandlerFunc(playbackAPI.GetPlaybackState).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var playbackState models.PlaybackState
	err = json.NewDecoder(rr.Body).Decode(&playbackState)
	assert.NoError(t, err)
	assert.Equal(t, "Chill Song 1", playbackState.CurrentSong.Title)
}
