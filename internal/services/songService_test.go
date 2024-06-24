package services

import (
	"github.com/stretchr/testify/assert"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"testing"
)

func TestCreateSong(t *testing.T) {
	storage := repositories.NewSongStorageMock()
	service := NewSongService(storage)

	song, err := service.CreateSong("Synth Beats", "Artist 1", "synth", "123", true, []string{"electronic", "beats"})
	assert.NoError(t, err)
	assert.NotNil(t, song)
	assert.Equal(t, "Synth Beats", song.Title)
	assert.Equal(t, "Artist 1", song.Artist)
	assert.ElementsMatch(t, []string{"electronic", "beats"}, []string{song.Tags[0].Name, song.Tags[1].Name})
}

func TestUpdateSong(t *testing.T) {
	storage := repositories.NewSongStorageMock()
	service := NewSongService(storage)

	song, err := service.CreateSong("Synth Beats", "Artist 1", "synth", "123", true, []string{"electronic", "beats"})
	assert.NoError(t, err)

	updatedSong, err := service.UpdateSong(&models.Song{
		ID:          song.ID,
		Title:       "Updated Beats",
		Artist:      "Artist 2",
		Genre:       "electronic",
		SunoID:      "456",
		IsGenerated: true,
	}, []string{"electronic", "updated"})
	assert.NoError(t, err)
	assert.NotNil(t, updatedSong)
	assert.Equal(t, "Updated Beats", updatedSong.Title)
	assert.Equal(t, "Artist 2", updatedSong.Artist)
	assert.ElementsMatch(t, []string{"electronic", "updated"}, []string{updatedSong.Tags[0].Name, updatedSong.Tags[1].Name})
}

func TestDeleteSong(t *testing.T) {
	storage := repositories.NewSongStorageMock()
	service := NewSongService(storage)

	song, err := service.CreateSong("Synth Beats", "Artist 1", "synth", "123", true, []string{"electronic", "beats"})
	assert.NoError(t, err)

	err = service.DeleteSong(song.ID)
	assert.NoError(t, err)

	_, err = service.GetSongByID(song.ID)
	assert.Error(t, err)
}

func TestGetSongByID(t *testing.T) {
	storage := repositories.NewSongStorageMock()
	service := NewSongService(storage)

	song, err := service.CreateSong("Synth Beats", "Artist 1", "synth", "123", true, []string{"electronic", "beats"})
	assert.NoError(t, err)

	fetchedSong, err := service.GetSongByID(song.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedSong)
	assert.Equal(t, "Synth Beats", fetchedSong.Title)
	assert.Equal(t, "Artist 1", fetchedSong.Artist)
	assert.ElementsMatch(t, []string{"electronic", "beats"}, []string{fetchedSong.Tags[0].Name, fetchedSong.Tags[1].Name})
}

func TestGetSongBySunoID(t *testing.T) {
	storage := repositories.NewSongStorageMock()
	service := NewSongService(storage)

	_, err := service.CreateSong("Synth Beats", "Artist 1", "synth", "123", true, []string{"electronic", "beats"})
	assert.NoError(t, err)

	fetchedSong, err := service.GetSongBySunoID("123")
	assert.NoError(t, err)
	assert.NotNil(t, fetchedSong)
	assert.Equal(t, "Synth Beats", fetchedSong.Title)
	assert.Equal(t, "Artist 1", fetchedSong.Artist)
	assert.ElementsMatch(t, []string{"electronic", "beats"}, []string{fetchedSong.Tags[0].Name, fetchedSong.Tags[1].Name})
}

func TestGetAllSongs(t *testing.T) {
	storage := repositories.NewSongStorageMock()
	service := NewSongService(storage)

	_, err := service.CreateSong("Synth Beats", "Artist 1", "synth", "123", true, []string{"electronic", "beats"})
	assert.NoError(t, err)
	_, err = service.CreateSong("Lo-fi Chill", "Artist 2", "lofi", "456", true, []string{"chill", "lofi"})
	assert.NoError(t, err)

	songs, err := service.GetAllSongs()
	assert.NoError(t, err)
	assert.Len(t, songs, 2)
}
