// Intializes the application backend and starts the server
package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"louderspace/config"
	"louderspace/internal/api"
	"louderspace/internal/repositories"
	"louderspace/internal/services"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	userStorage := repositories.NewUserDatabase(db)
	songStorage := repositories.NewSongDatabase(db)
	stationStorage := repositories.NewStationDatabase(db)

	userService := services.NewUserService(userStorage)
	stationService := services.NewStationService(stationStorage)
	playbackService := services.NewPlaybackService(stationStorage)
	songService := services.NewSongService(songStorage)

	userAPI := api.NewUserAPI(userService)
	stationAPI := api.NewStationAPI(stationService)
	playbackAPI := api.NewPlaybackAPI(playbackService)
	songAPI := api.NewSongAPI(songService)

	http.HandleFunc("/register", userAPI.Register)
	http.HandleFunc("/login", userAPI.Login)

	http.HandleFunc("/stations/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/stations/")
		if strings.HasSuffix(path, "/songs") {
			stationIDStr := strings.TrimSuffix(path, "/songs")
			if stationIDStr != "" {
				stationID, err := strconv.Atoi(stationIDStr)
				if err != nil {
					http.Error(w, "Invalid station ID", http.StatusBadRequest)
					return
				}
				// Call GetSongsForStation with the station ID
				stationAPI.GetSongsForStationByID(w, r, stationID)
			} else {
				http.NotFound(w, r)
			}
		} else {
			switch r.Method {
			case "POST":
				stationAPI.CreateStation(w, r)
			case "PUT":
				stationAPI.UpdateStation(w, r)
			case "DELETE":
				stationAPI.DeleteStation(w, r)
			case "GET":
				stationAPI.GetAllStations(w, r)
			}
		}
	})
	http.HandleFunc("/playback/play", playbackAPI.Play)
	http.HandleFunc("/playback/pause", playbackAPI.Pause)
	http.HandleFunc("/playback/skip", playbackAPI.Skip)
	http.HandleFunc("/playback/rewind", playbackAPI.Rewind)
	http.HandleFunc("/playback/state", playbackAPI.GetPlaybackState)

	http.HandleFunc("/songs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			songAPI.CreateSong(w, r)
		case "GET":
			songAPI.GetSong(w, r)
		}
	})
	http.HandleFunc("/songs/suno", songAPI.GetSongBySunoID)

	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
