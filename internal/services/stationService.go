package services

import (
	"log"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
)

type StationManagement interface {
	CreateStation(name string, tags []string) (*models.Station, error)
	UpdateStation(id int, name string, tags []string) (*models.Station, error)
	DeleteStation(id int) error
	GetStation(id int) (*models.Station, error)
	GetAllStations() ([]*models.Station, error)
	GetSongsForStation(stationID int) ([]*models.Song, error)
	GetSongsForStationWithFeedback(stationID, userID int) ([]*models.SongWithFeedback, error)
}

type StationService struct {
	stationStorage  repositories.StationStorage
	feedbackStorage repositories.FeedbackStorage
	songStorage     repositories.SongStorage
}

func NewStationService(stationStorage repositories.StationStorage, feedbackStorage repositories.FeedbackStorage, songStorage repositories.SongStorage) StationManagement {
	return &StationService{stationStorage, feedbackStorage, songStorage}
}

func (s *StationService) CreateStation(name string, tags []string) (*models.Station, error) {
	station := &models.Station{Name: name, Tags: tags}
	if err := s.stationStorage.Create(station); err != nil {
		log.Printf("Error creating station: %v", err)
		return nil, err
	}
	log.Printf("Station created: %v", station)
	return station, nil
}

func (s *StationService) UpdateStation(id int, name string, tags []string) (*models.Station, error) {
	station := &models.Station{ID: id, Name: name, Tags: tags}
	if err := s.stationStorage.Update(station); err != nil {
		return nil, err
	}
	return station, nil
}

func (s *StationService) DeleteStation(id int) error {
	return s.stationStorage.Delete(id)
}

func (s *StationService) GetStation(id int) (*models.Station, error) {
	return s.stationStorage.ByID(id)
}

func (s *StationService) GetAllStations() ([]*models.Station, error) {
	return s.stationStorage.All()
}

func (s *StationService) GetSongsForStation(stationID int) ([]*models.Song, error) {
	station, err := s.stationStorage.ByID(stationID)
	if err != nil {
		return nil, err
	}
	return s.stationStorage.SongsByTags(station.Tags)
}

func (s *StationService) GetSongsForStationWithFeedback(stationID, userID int) ([]*models.SongWithFeedback, error) {
	songs, err := s.songStorage.ByStationID(stationID)
	if err != nil {
		return nil, err
	}

	// Get song IDs
	songIDs := make([]int, len(songs))
	for i, song := range songs {
		songIDs[i] = song.ID
	}

	// Get feedback for the songs
	feedbackMap, err := s.feedbackStorage.GetFeedbackForUserAndSongs(userID, songIDs)
	if err != nil {
		return nil, err
	}

	// Combine song and feedback information
	songsWithFeedback := make([]*models.SongWithFeedback, len(songs))
	for i, song := range songs {
		liked, ok := feedbackMap[song.ID]
		if !ok {
			liked = false
		}
		songsWithFeedback[i] = &models.SongWithFeedback{
			Song:  song,
			Liked: liked,
		}
	}

	return songsWithFeedback, nil
}
