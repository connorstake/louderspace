package main

import (
	"database/sql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"louderspace/config"
	"louderspace/internal/api"
	"louderspace/internal/logger"
	"louderspace/internal/middleware"
	"louderspace/internal/models"
	"louderspace/internal/repositories"
	"louderspace/internal/services"
	"net/http"
)

func main() {
	logger.Init()

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
	playEventStorage := repositories.NewPlayEventDatabase(db)
	feedbackStorage := repositories.NewFeedbackDatabase(db)
	pomodoroSessionStorage := repositories.NewPomodoroSessionDatabase(db)

	userService := services.NewUserService(userStorage)
	stationService := services.NewStationService(stationStorage, feedbackStorage, songStorage)
	playbackService := services.NewPlaybackService(stationStorage)
	songService := services.NewSongService(songStorage)
	tagService := services.NewTagService(tagStorage)
	playEventService := services.NewPlayEventService(playEventStorage)
	feedbackService := services.NewFeedbackService(feedbackStorage)
	pomodoroSessionService := services.NewPomodoroSessionService(pomodoroSessionStorage)

	userAPI := api.NewUserAPI(userService)
	authAPI := api.NewAuthAPI(userService)
	stationAPI := api.NewStationAPI(stationService)
	playbackAPI := api.NewPlaybackAPI(playbackService)
	songAPI := api.NewSongAPI(songService)
	tagAPI := api.NewTagAPI(tagService)
	playEventAPI := api.NewPlayEventAPI(playEventService)
	feedbackAPI := api.NewFeedbackAPI(feedbackService)
	pomodoroAPI := api.NewPomodoroSessionAPI(pomodoroSessionService)

	r := mux.NewRouter()

	// Use logging middleware
	r.Use(middleware.LoggingMiddleware)

	// Public routes
	r.HandleFunc("/register", authAPI.Register).Methods("POST")
	r.HandleFunc("/login", authAPI.Login).Methods("POST")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.WithUser)

	protected.HandleFunc("/feedback", feedbackAPI.SaveFeedback).Methods("POST")
	protected.HandleFunc("/feedback", feedbackAPI.DeleteFeedback).Methods("DELETE")
	protected.HandleFunc("/feedback", feedbackAPI.GetFeedback).Methods("GET")

	protected.HandleFunc("/me", http.HandlerFunc(authAPI.Me)).Methods("GET")

	protected.HandleFunc("/stations", stationAPI.GetAllStations).Methods("GET")
	protected.HandleFunc("/stations/{id:[0-9]+}/songs", stationAPI.GetSongsForStationByID).Methods("GET")

	protected.HandleFunc("/playback/play", playbackAPI.Play).Methods("POST")
	protected.HandleFunc("/playback/pause", playbackAPI.Pause).Methods("POST")
	protected.HandleFunc("/playback/skip", playbackAPI.Skip).Methods("POST")
	protected.HandleFunc("/playback/rewind", playbackAPI.Rewind).Methods("POST")
	protected.HandleFunc("/playback/state", playbackAPI.GetPlaybackState).Methods("GET")

	protected.HandleFunc("/pomodoro/start", pomodoroAPI.StartSession).Methods("POST")
	protected.HandleFunc("/pomodoro/end", pomodoroAPI.EndSession).Methods("POST")

	protected.HandleFunc("/songs", songAPI.GetAllSongs).Methods("GET")

	adminRouter := protected.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.RequireRole(models.RoleAdmin))

	adminRouter.HandleFunc("/pomodoro/user/{user_id}/sessions", pomodoroAPI.GetSessionsByUser).Methods("GET")
	adminRouter.HandleFunc("/pomodoro/user/{user_id}/metrics", pomodoroAPI.GetFocusMetrics).Methods("GET")

	adminRouter.HandleFunc("/users", userAPI.Users).Methods("GET")

	adminRouter.HandleFunc("/play_events", playEventAPI.LogPlayEvent).Methods("POST")
	adminRouter.HandleFunc("/update_aggregates", playEventAPI.UpdateAggregates).Methods("POST")

	adminRouter.HandleFunc("/tags", tagAPI.CreateTag).Methods("POST")
	adminRouter.HandleFunc("/tags", tagAPI.GetTags).Methods("GET")
	adminRouter.HandleFunc("/tags/{id:[0-9]+}", tagAPI.UpdateTag).Methods("PUT")
	adminRouter.HandleFunc("/tags/{id:[0-9]+}", tagAPI.DeleteTag).Methods("DELETE")

	adminRouter.HandleFunc("/songs", songAPI.CreateSong).Methods("POST")
	adminRouter.HandleFunc("/songs/{id:[0-9]+}", songAPI.DeleteSong).Methods("DELETE")
	adminRouter.HandleFunc("/songs/{id:[0-9]+}", songAPI.UpdateSong).Methods("PUT")
	adminRouter.HandleFunc("/songs/{id:[0-9]+}", songAPI.GetSong).Methods("GET")
	adminRouter.HandleFunc("/songs/suno", songAPI.GetSongBySunoID).Methods("GET")

	adminRouter.HandleFunc("/stations", stationAPI.CreateStation).Methods("POST")
	adminRouter.HandleFunc("/stations/{id:[0-9]+}", stationAPI.UpdateStation).Methods("PUT")
	adminRouter.HandleFunc("/stations/{id:[0-9]+}", stationAPI.DeleteStation).Methods("DELETE")

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(r)))
}
