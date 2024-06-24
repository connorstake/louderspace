package main

import (
	"database/sql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"louderspace/config"
	"louderspace/internal/api"
	"louderspace/internal/repositories"
	"louderspace/internal/services"
	"net/http"
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

	r := mux.NewRouter()

	r.HandleFunc("/register", userAPI.Register).Methods("POST")
	r.HandleFunc("/login", userAPI.Login).Methods("POST")
	r.HandleFunc("/users", userAPI.Users).Methods("GET")

	r.HandleFunc("/stations", stationAPI.CreateStation).Methods("POST")
	r.HandleFunc("/stations", stationAPI.GetAllStations).Methods("GET")
	r.HandleFunc("/stations/{id:[0-9]+}", stationAPI.UpdateStation).Methods("PUT")
	r.HandleFunc("/stations/{id:[0-9]+}/songs", stationAPI.GetSongsForStationByID).Methods("GET")
	r.HandleFunc("/stations/{id:[0-9]+}", stationAPI.DeleteStation).Methods("DELETE")

	r.HandleFunc("/playback/play", playbackAPI.Play).Methods("POST")
	r.HandleFunc("/playback/pause", playbackAPI.Pause).Methods("POST")
	r.HandleFunc("/playback/skip", playbackAPI.Skip).Methods("POST")
	r.HandleFunc("/playback/rewind", playbackAPI.Rewind).Methods("POST")
	r.HandleFunc("/playback/state", playbackAPI.GetPlaybackState).Methods("GET")

	r.HandleFunc("/songs", songAPI.CreateSong).Methods("POST")
	r.HandleFunc("/songs", songAPI.GetAllSongs).Methods("GET")
	r.HandleFunc("/songs", songAPI.UpdateSong).Methods("PUT")
	r.HandleFunc("/songs", songAPI.DeleteSong).Methods("DELETE")
	r.HandleFunc("/songs/{id:[0-9]+}", songAPI.GetSong).Methods("GET")
	r.HandleFunc("/songs/suno", songAPI.GetSongBySunoID).Methods("GET")

	r.HandleFunc("/tags", tagAPI.GetTags).Methods("GET")
	r.HandleFunc("/tags", tagAPI.CreateTag).Methods("POST")
	r.HandleFunc("/tags/{id:[0-9]+}", tagAPI.UpdateTag).Methods("PUT")
	r.HandleFunc("/tags/{id:[0-9]+}", tagAPI.DeleteTag).Methods("DELETE")

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Adjust the allowed origins as needed
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(r)))
}
