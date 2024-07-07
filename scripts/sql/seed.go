package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	host, _ := os.LookupEnv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	for _, value := range os.Environ() {
		fmt.Println(value)
	}

	fmt.Printf("host=%s port=%d user=%s ", host, port, user)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	// Truncate tables
	truncateTables(db)

	// Seed data
	err = seedUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	err = seedTags(db)
	if err != nil {
		log.Fatal(err)
	}

	err = seedSongs(db)
	if err != nil {
		log.Fatal(err)
	}

	err = seedSongTags(db)
	if err != nil {
		log.Fatal(err)
	}

	err = seedStations(db)
	if err != nil {
		log.Fatal(err)
	}
}

func truncateTables(db *sql.DB) {
	tables := []string{"song_tags", "songs", "users", "tags", "stations"}
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", table))
		if err != nil {
			log.Fatalf("Failed to truncate table %s: %v", table, err)
		}
	}
}

func seedUsers(db *sql.DB) error {
	users := []struct {
		username string
		password string
		email    string
		role     string
	}{
		{"user1", "password1", "user1@example.com", "admin"},
		{"user2", "password2", "user2@example.com", "free"},
	}

	for _, u := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %v", err)
		}

		_, err = db.Exec("INSERT INTO users (username, password, email, role) VALUES ($1, $2, $3, $4)",
			u.username, hashedPassword, u.email, u.role)
		if err != nil {
			return fmt.Errorf("failed to insert user %s: %v", u.username, err)
		}
	}

	return nil
}

func seedTags(db *sql.DB) error {
	tags := []string{
		"chill", "beats", "vibes", "lofi", "hiphop", "synth", "instrumental", "classical", "piano",
	}

	for _, tag := range tags {
		_, err := db.Exec("INSERT INTO tags (name) VALUES ($1)", tag)
		if err != nil {
			return fmt.Errorf("failed to insert tag %s: %v", tag, err)
		}
	}

	return nil
}

func seedSongs(db *sql.DB) error {
	songs := []struct {
		title        string
		artist       string
		genre        []string
		suno_id      string
		is_generated bool
	}{
		{"Synth Beats 1", "Artist 4", []string{"synth", "instrumental", "beats", "lofi", "hiphop"}, "6e3cd1cf-f487-487f-b53b-e858b9b101eb", true},
		{"Instrumental Vibes", "Artist 5", []string{"synth", "instrumental", "beats", "lofi", "hiphop"}, "eee20928-b7ab-44f5-8a9f-31f2f7d568a5", true},
		{"Classical Hiphop", "Artist 6", []string{"synth", "instrumental", "beats", "lofi", "hiphop"}, "6d436627-10c0-4a02-a384-5d13a27d8d96", true},
		{"Piano Lo-fi", "Artist 7", []string{"synth", "instrumental", "beats", "lofi", "hiphop"}, "3c6534d5-fab4-4e9a-b230-471a76debcbf", true},
		{"Synthwave Chill", "Artist 8", []string{"synth", "instrumental", "beats", "lofi", "hiphop"}, "7be45898-26a1-479b-bbbd-aaa39ec83551", true},
		{"Instrumental Beats", "Artist 9", []string{"synth", "instrumental", "beats", "lofi", "hiphop"}, "b089275f-f73f-4094-b5ae-77e98b1c3311", true},
		{"Song 1", "Random Artist", []string{"classical", "lofi"}, "6033d8c5-e024-4d84-80a1-1df28683b304", true},
		{"Song 2", "Random Artist", []string{"classical", "lofi"}, "c7b04693-164a-448e-a98b-1ce02bf750a1", true},
		{"Song 3", "Random Artist", []string{"classical", "lofi"}, "6f3d4e9f-90a7-4e87-b6bd-5d0b67085b25", true},
		{"Song 4", "Random Artist", []string{"classical", "lofi"}, "e37d4828-44f4-4394-93fe-fcb4eee6619e", true},
		{"Song 5", "Random Artist", []string{"classical", "lofi"}, "d8c2a782-5d34-455d-8f67-f030172a3e69", true},
		{"Song 6", "Random Artist", []string{"classical", "lofi"}, "67a74714-bc38-4247-b936-a2bf29f7854f", true},
		{"Song 7", "Random Artist", []string{"classical", "lofi"}, "96edec8b-97af-47ab-a1ee-0cb692725d4f", true},
		{"Song 8", "Random Artist", []string{"classical", "lofi"}, "b0959337-1f4c-447d-88b6-32bf50bbfa90", true},
		{"Song 9", "Random Artist", []string{"classical", "lofi"}, "6e4b0a38-cbe3-469f-98d4-74ef9099a2b2", true},
		{"Song 10", "Random Artist", []string{"classical", "lofi"}, "3073ec18-bb02-4a15-b8a8-223680e83bd8", true},
		{"Song 11", "Random Artist", []string{"classical", "lofi"}, "47b44acb-9432-4aef-b26d-1e1b665729a1", true},
		{"Song 12", "Random Artist", []string{"classical", "lofi"}, "adc3ca43-cb55-41c0-ae7f-95a002690939", true},
		{"Song 13", "Random Artist", []string{"classical", "lofi"}, "29a278fc-1ad5-4bdb-98a3-50cff1c60ae4", true},
		{"Song 14", "Random Artist", []string{"classical", "lofi"}, "ed4bfcc3-f68b-41e1-8199-73017b7c44f3", true},
		{"Song 15", "Random Artist", []string{"classical", "lofi"}, "1d5128b8-ee1d-4b9b-804a-09ef79da5ece", true},
		{"Song 16", "Random Artist", []string{"classical", "lofi"}, "ae4094c1-36aa-489b-a046-f8b2ef9da914", true},
	}

	for _, s := range songs {
		_, err := db.Exec("INSERT INTO songs (title, artist, genre, suno_id, is_generated) VALUES ($1, $2, $3, $4, $5)",
			s.title, s.artist, pq.Array(s.genre), s.suno_id, s.is_generated)
		if err != nil {
			return fmt.Errorf("failed to insert song %s: %v", s.title, err)
		}
	}

	return nil
}

func seedSongTags(db *sql.DB) error {
	songTags := []struct {
		song_id int
		tag_id  int
	}{
		{1, 6}, {1, 7}, {1, 2}, {1, 4}, {1, 5},
		{2, 6}, {2, 7}, {2, 2}, {2, 4}, {2, 5},
		{3, 6}, {3, 7}, {3, 2}, {3, 4}, {3, 5},
		{4, 6}, {4, 7}, {4, 2}, {4, 4}, {4, 5},
		{5, 6}, {5, 7}, {5, 2}, {5, 4}, {5, 5},
		{6, 6}, {6, 7}, {6, 2}, {6, 4}, {6, 5},
		{7, 8}, {7, 4},
		{8, 8}, {8, 4},
		{9, 8}, {9, 4},
		{10, 8}, {10, 4},
		{11, 8}, {11, 4},
		{12, 8}, {12, 4},
		{13, 8}, {13, 4},
		{14, 8}, {14, 4},
		{15, 8}, {15, 4},
		{16, 8}, {16, 4},
		{17, 8}, {17, 4},
		{18, 8}, {18, 4},
		{19, 8}, {19, 4},
		{20, 8}, {20, 4},
		{21, 8}, {21, 4},
		{22, 8}, {22, 4},
	}

	for _, st := range songTags {
		_, err := db.Exec("INSERT INTO song_tags (song_id, tag_id) VALUES ($1, $2)", st.song_id, st.tag_id)
		if err != nil {
			return fmt.Errorf("failed to insert song_tag (%d, %d): %v", st.song_id, st.tag_id, err)
		}
	}

	return nil
}

func seedStations(db *sql.DB) error {
	stations := []struct {
		name string
		tags string
	}{
		{"Coding Den", "synth, lofi"},
		{"Reading Room", "classical, lofi"},
	}

	for _, s := range stations {
		_, err := db.Exec("INSERT INTO stations (name, tags) VALUES ($1, $2)", s.name, s.tags)
		if err != nil {
			return fmt.Errorf("failed to insert station %s: %v", s.name, err)
		}
	}

	return nil
}
