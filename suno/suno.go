package suno

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

// Request struct for creating a song
type CreateSongRequest struct {
	Prompt         string `json:"prompt"`
	Tags           string `json:"tags,omitempty"`
	CustomMode     bool   `json:"custom_mode"`
	Title          string `json:"title,omitempty"`
	ContinueAt     int    `json:"continue_at,omitempty"`
	ContinueClipID string `json:"continue_clip_id,omitempty"`
}

// Response struct for creating a song
type CreateSongResponse struct {
	Code int `json:"code"`
	Data struct {
		TaskID string `json:"task_id"`
	} `json:"data"`
	Message string `json:"message"`
}

// Error response struct
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Function to generate a song using the Suno API
func CreateSong(request CreateSongRequest) (*CreateSongResponse, error) {
	// Suno API URL
	url := "https://api.sunoapi.com/api/v1/suno/create"

	// Get the API token from environment variables
	token := os.Getenv("SUNO_API_TOKEN")
	if token == "" {
		return nil, errors.New("SUNO_API_TOKEN is not set")
	}

	// Marshal the request body
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create an HTTP client and set a timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status is not OK
	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}
		return nil, errors.New(errorResponse.Message)
	}

	// Decode the response body
	var createSongResponse CreateSongResponse
	if err := json.NewDecoder(resp.Body).Decode(&createSongResponse); err != nil {
		return nil, err
	}

	return &createSongResponse, nil
}

// Response struct for retrieving generated music
type GetClipResponse struct {
	Code int `json:"code"`
	Data struct {
		TaskID   string          `json:"task_id"`
		Status   string          `json:"status"`
		Input    string          `json:"input"`
		Clips    map[string]Clip `json:"clips"`
		Metadata Metadata        `json:"metadata"`
	} `json:"data"`
	Message string `json:"message"`
}

// Clip struct representing each generated clip
type Clip struct {
	ID                string      `json:"id"`
	VideoURL          string      `json:"video_url"`
	AudioURL          string      `json:"audio_url"`
	ImageURL          string      `json:"image_url"`
	ImageLargeURL     string      `json:"image_large_url"`
	IsVideoPending    bool        `json:"is_video_pending"`
	MajorModelVersion string      `json:"major_model_version"`
	ModelName         string      `json:"model_name"`
	Metadata          Metadata    `json:"metadata"`
	IsLiked           bool        `json:"is_liked"`
	UserID            string      `json:"user_id"`
	DisplayName       string      `json:"display_name"`
	Handle            string      `json:"handle"`
	IsHandleUpdated   bool        `json:"is_handle_updated"`
	IsTrashed         bool        `json:"is_trashed"`
	Reaction          interface{} `json:"reaction"`
	CreatedAt         string      `json:"created_at"`
	Status            string      `json:"status"`
	Title             string      `json:"title"`
	PlayCount         int         `json:"play_count"`
	UpvoteCount       int         `json:"upvote_count"`
	IsPublic          bool        `json:"is_public"`
}

// Metadata struct for additional details of the clip
type Metadata struct {
	Tags                 string      `json:"tags"`
	Prompt               string      `json:"prompt"`
	GptDescriptionPrompt string      `json:"gpt_description_prompt"`
	AudioPromptID        string      `json:"audio_prompt_id"`
	History              interface{} `json:"history"`
	ConcatHistory        interface{} `json:"concat_history"`
	Type                 string      `json:"type"`
	Duration             float64     `json:"duration"`
	RefundCredits        bool        `json:"refund_credits"`
	Stream               bool        `json:"stream"`
	ErrorType            string      `json:"error_type"`
	ErrorMessage         string      `json:"error_message"`
}

// Function to get generated music using the task ID
func GetClip(taskID string) (*GetClipResponse, error) {
	url := "https://api.sunoapi.com/api/v1/suno/clip/" + taskID
	token := os.Getenv("SUNO_API_TOKEN")
	if token == "" {
		return nil, errors.New("SUNO_API_TOKEN is not set")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}
		return nil, errors.New(errorResponse.Message)
	}

	var getClipResponse GetClipResponse
	if err := json.NewDecoder(resp.Body).Decode(&getClipResponse); err != nil {
		return nil, err
	}

	return &getClipResponse, nil
}
