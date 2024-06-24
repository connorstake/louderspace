package services

import (
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"time"
)

type SongManagement interface {
	CreateSong(title, artist, genre, sunoID string, isGenerated bool) (*models.Song, error)
	GetSongByID(songID int) (*models.Song, error)
	GetSongBySunoID(sunoID string) (*models.Song, error)
	GetAllSongs() ([]*models.Song, error)
	GetSongsByTags(tags []string) ([]*models.Song, error)
	UpdateSong(id int, title, artist, genre, sunoID string, isGenerated bool) (*models.Song, error)
	DeleteSong(songID int) error
	GetSongsForStation(stationID int) ([]*models.Song, error)
}

type SongService struct {
	songStorage    repositories.SongStorage
	stationStorage repositories.StationStorage
}

func NewSongService(songStorage repositories.SongStorage, stationStorage repositories.StationStorage) SongManagement {
	return &SongService{songStorage, stationStorage}
}

func (s *SongService) CreateSong(title, artist, genre, sunoID string, isGenerated bool) (*models.Song, error) {
	song := &models.Song{
		Title:       title,
		Artist:      artist,
		Genre:       genre,
		SunoID:      sunoID,
		IsGenerated: isGenerated,
		CreatedAt:   time.Now(),
	}
	if err := s.songStorage.Create(song); err != nil {
		return nil, err
	}
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

func (s *SongService) GetSongsByTags(tags []string) ([]*models.Song, error) {
	return s.songStorage.SongsByTags(tags)
}

func (s *SongService) UpdateSong(id int, title, artist, genre, sunoID string, isGenerated bool) (*models.Song, error) {
	song, err := s.songStorage.ByID(id)
	if err != nil {
		return nil, err
	}

	song.Title = title
	song.Artist = artist
	song.Genre = genre
	song.SunoID = sunoID
	song.IsGenerated = isGenerated

	if err := s.songStorage.Update(song); err != nil {
		return nil, err
	}
	return song, nil
}

func (s *SongService) DeleteSong(songID int) error {
	return s.songStorage.Delete(songID)
}

func (s *SongService) GetSongsForStation(stationID int) ([]*models.Song, error) {
	station, err := s.stationStorage.ByID(stationID)
	if err != nil {
		return nil, err
	}

	return s.songStorage.SongsByTags(station.Tags)
}
