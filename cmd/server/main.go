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

// CORSMiddleware sets the CORS headers for incoming requests.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	mux := http.NewServeMux()

	userStorage := repositories.NewUserDatabase(db)
	songStorage := repositories.NewSongDatabase(db)
	stationStorage := repositories.NewStationDatabase(db)
	tagStorage := repositories.NewTagDatabase(db)

	userService := services.NewUserService(userStorage)
	stationService := services.NewStationService(stationStorage)
	playbackService := services.NewPlaybackService(stationStorage)
	songService := services.NewSongService(songStorage, stationStorage)
	tagService := services.NewTagService(tagStorage)

	userAPI := api.NewUserAPI(userService)
	stationAPI := api.NewStationAPI(stationService)
	playbackAPI := api.NewPlaybackAPI(playbackService)
	songAPI := api.NewSongAPI(songService)
	tagAPI := api.NewTagAPI(tagService)

	mux.HandleFunc("/register", userAPI.Register)
	mux.HandleFunc("/login", userAPI.Login)
	mux.HandleFunc("/users", userAPI.Users)

	mux.HandleFunc("/stations", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/stations")
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
	mux.HandleFunc("/playback/play", playbackAPI.Play)
	mux.HandleFunc("/playback/pause", playbackAPI.Pause)
	mux.HandleFunc("/playback/skip", playbackAPI.Skip)
	mux.HandleFunc("/playback/rewind", playbackAPI.Rewind)
	mux.HandleFunc("/playback/state", playbackAPI.GetPlaybackState)

	mux.HandleFunc("/songs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			songAPI.CreateSong(w, r)
		case "GET":
			idStr := r.URL.Query().Get("id")
			if idStr == "" {
				songAPI.GetAllSongs(w, r)
			} else {
				songAPI.GetSong(w, r)
			}
		case "PUT":
			songAPI.UpdateSong(w, r)
		case "DELETE":
			songAPI.DeleteSong(w, r)
		}
	})
	mux.HandleFunc("/songs/suno", songAPI.GetSongBySunoID)

	mux.HandleFunc("/tags", tagAPI.GetTags)

	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", CORSMiddleware(mux)))

}
