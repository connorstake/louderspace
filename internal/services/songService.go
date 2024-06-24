package services

import (
	"louderspace/internal/logger"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"time"
)

type SongManagement interface {
	CreateSong(title, artist, genre, sunoID string, isGenerated bool, tags []string) (*models.Song, error)
	GetSongByID(songID int) (*models.Song, error)
	GetSongBySunoID(sunoID string) (*models.Song, error)
	GetAllSongs() ([]*models.Song, error)
	GetSongsForStation(stationID int) ([]*models.Song, error)
	UpdateSong(song *models.Song, tags []string) (*models.Song, error)
	DeleteSong(id int) error
}

type SongService struct {
	songStorage repositories.SongStorage
}

func NewSongService(songStorage repositories.SongStorage) SongManagement {
	return &SongService{songStorage}
}

func (s *SongService) CreateSong(title, artist, genre, sunoID string, isGenerated bool, tags []string) (*models.Song, error) {
	song := &models.Song{
		Title:       title,
		Artist:      artist,
		Genre:       genre,
		SunoID:      sunoID,
		IsGenerated: isGenerated,
		CreatedAt:   time.Now(),
	}
	if err := s.songStorage.Create(song, tags); err != nil {
		logger.Error("Failed to create song:", err)
		return nil, err
	}
	// Fetch updated tags
	updatedTags, err := s.songStorage.GetTagsBySongID(song.ID)
	if err != nil {
		logger.Error("Failed to fetch updated tags:", err)
		return nil, err
	}

	song.Tags = updatedTags
	return song, nil
}

func (s *SongService) UpdateSong(song *models.Song, tags []string) (*models.Song, error) {

	if err := s.songStorage.Update(song, tags); err != nil {
		logger.Error("Failed to update song:", err)
		return nil, err
	}

	// Fetch updated tags
	updatedTags, err := s.songStorage.GetTagsBySongID(song.ID)
	if err != nil {
		logger.Error("Failed to fetch updated tags:", err)
		return nil, err
	}

	song.Tags = updatedTags
	return song, nil
}

func (s *SongService) GetSongByID(songID int) (*models.Song, error) {
	return s.songStorage.ByID(songID)
}

func (s *SongService) GetSongBySunoID(sunoID string) (*models.Song, error) {
	return s.songStorage.BySunoID(sunoID)
}

func (s *SongService) GetAllSongs() ([]*models.Song, error) {
	return s.songStorage.All()
}

func (s *SongService) GetSongsForStation(stationID int) ([]*models.Song, error) {
	return s.songStorage.ByStationID(stationID)
}

func (s *SongService) DeleteSong(id int) error {
	return s.songStorage.Delete(id)
}
