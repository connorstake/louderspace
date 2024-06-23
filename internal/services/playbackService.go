package services

import (
	"errors"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"sync"
	"time"
)

type PlaybackManagement interface {
	Play(userID, stationID int) (*models.PlaybackState, error)
	Pause(userID int) (*models.PlaybackState, error)
	Skip(userID int) (*models.PlaybackState, error)
	Rewind(userID int) (*models.PlaybackState, error)
	GetPlaybackState(userID int) (*models.PlaybackState, error)
}

type PlaybackService struct {
	stationStorage repositories.StationStorage
	userPlayback   map[int]*models.PlaybackState
	mu             sync.Mutex
}

func NewPlaybackService(stationStorage repositories.StationStorage) PlaybackManagement {
	return &PlaybackService{
		stationStorage: stationStorage,
		userPlayback:   make(map[int]*models.PlaybackState),
	}
}

func (p *PlaybackService) Play(userID, stationID int) (*models.PlaybackState, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	playbackState, exists := p.userPlayback[userID]
	if !exists || playbackState.StationID != stationID {
		station, err := p.stationStorage.ByID(stationID)
		if err != nil {
			return nil, err
		}

		songs, err := p.stationStorage.SongsByTags(station.Tags)
		if err != nil {
			return nil, err
		}

		if len(songs) == 0 {
			return nil, errors.New("no songs found for this station")
		}

		playbackState = &models.PlaybackState{
			UserID:       userID,
			StationID:    stationID,
			CurrentSong:  songs[0],
			SongQueue:    songs[1:],
			IsPlaying:    true,
			PlaybackTime: time.Now(),
		}
		p.userPlayback[userID] = playbackState
	} else {
		playbackState.IsPlaying = true
	}

	return playbackState, nil
}

func (p *PlaybackService) Pause(userID int) (*models.PlaybackState, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	playbackState, exists := p.userPlayback[userID]
	if !exists {
		return nil, errors.New("no playback state found for this user")
	}

	playbackState.IsPlaying = false
	return playbackState, nil
}

func (p *PlaybackService) Skip(userID int) (*models.PlaybackState, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	playbackState, exists := p.userPlayback[userID]
	if !exists {
		return nil, errors.New("no playback state found for this user")
	}

	if len(playbackState.SongQueue) == 0 {
		return nil, errors.New("no more songs in the queue")
	}

	playbackState.CurrentSong = playbackState.SongQueue[0]
	playbackState.SongQueue = playbackState.SongQueue[1:]
	playbackState.PlaybackTime = time.Now()
	playbackState.IsPlaying = true

	return playbackState, nil
}

func (p *PlaybackService) Rewind(userID int) (*models.PlaybackState, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	playbackState, exists := p.userPlayback[userID]
	if !exists {
		return nil, errors.New("no playback state found for this user")
	}

	if playbackState.CurrentSong == nil {
		return nil, errors.New("no current song to rewind")
	}

	playbackState.PlaybackTime = time.Now()
	return playbackState, nil
}

func (p *PlaybackService) GetPlaybackState(userID int) (*models.PlaybackState, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	playbackState, exists := p.userPlayback[userID]
	if !exists {
		return nil, errors.New("no playback state found for this user")
	}

	return playbackState, nil
}
