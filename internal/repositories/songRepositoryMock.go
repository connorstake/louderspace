package repositories

import (
	"errors"
	"louderspace/internal/models"
	"sync"
)

type SongStorageMock struct {
	songs  map[int]*models.Song
	tags   map[int][]models.Tag
	nextID int
	mu     sync.RWMutex
}

func NewSongStorageMock() *SongStorageMock {
	return &SongStorageMock{
		songs:  make(map[int]*models.Song),
		tags:   make(map[int][]models.Tag),
		nextID: 1,
	}
}

func (s *SongStorageMock) Create(song *models.Song, tags []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	song.ID = s.nextID
	s.nextID++
	s.songs[song.ID] = song

	for _, tag := range tags {
		tagID := s.nextID
		s.tags[song.ID] = append(s.tags[song.ID], models.Tag{ID: tagID, Name: tag})
		s.nextID++
	}
	return nil
}

func (s *SongStorageMock) Update(song *models.Song, tags []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.songs[song.ID]; !exists {
		return errors.New("song not found")
	}

	s.songs[song.ID] = song
	s.tags[song.ID] = []models.Tag{}

	for _, tag := range tags {
		tagID := s.nextID
		s.tags[song.ID] = append(s.tags[song.ID], models.Tag{ID: tagID, Name: tag})
		s.nextID++
	}
	return nil
}

func (s *SongStorageMock) ByID(id int) (*models.Song, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	song, exists := s.songs[id]
	if !exists {
		return nil, errors.New("song not found")
	}

	song.Tags = s.tags[id]
	return song, nil
}

func (s *SongStorageMock) BySunoID(sunoID string) (*models.Song, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, song := range s.songs {
		if song.SunoID == sunoID {
			song.Tags = s.tags[song.ID]
			return song, nil
		}
	}
	return nil, errors.New("song not found")
}

func (s *SongStorageMock) All() ([]*models.Song, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var songs []*models.Song
	for _, song := range s.songs {
		song.Tags = s.tags[song.ID]
		songs = append(songs, song)
	}
	return songs, nil
}

func (s *SongStorageMock) ByStationID(stationID int) ([]*models.Song, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var songs []*models.Song
	for _, song := range s.songs {
		for _, tag := range song.Tags {
			if containsTag(s.tags[song.ID], tag.Name) {
				songs = append(songs, song)
				break
			}
		}
	}
	return songs, nil
}

func (s *SongStorageMock) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.songs[id]; !exists {
		return errors.New("song not found")
	}

	delete(s.songs, id)
	delete(s.tags, id)
	return nil
}

func (s *SongStorageMock) GetTagsBySongID(songID int) ([]models.Tag, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tags, exists := s.tags[songID]
	if !exists {
		return nil, errors.New("no tags found for song")
	}

	return tags, nil
}

func containsTag(tags []models.Tag, tagName string) bool {
	for _, tag := range tags {
		if tag.Name == tagName {
			return true
		}
	}
	return false
}
