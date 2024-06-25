CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(50) NOT NULL DEFAULT 'free'
    );

CREATE TABLE IF NOT EXISTS songs (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR(100) NOT NULL,
    artist VARCHAR(100),
    genre VARCHAR(50),
    suno_id VARCHAR(50) UNIQUE NOT NULL,
    is_generated BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS tags (
                                    id SERIAL PRIMARY KEY,
                                    name VARCHAR(50) UNIQUE NOT NULL
    );

CREATE TABLE IF NOT EXISTS song_tags (
                                         id SERIAL PRIMARY KEY,
                                         song_id INT NOT NULL REFERENCES songs(id),
    tag_id INT NOT NULL REFERENCES tags(id)
    );

CREATE TABLE IF NOT EXISTS feedback (
                                        id SERIAL PRIMARY KEY,
                                        user_id INT NOT NULL REFERENCES users(id),
    song_id INT NOT NULL REFERENCES songs(id),
    liked BOOLEAN NOT NULL
    );

CREATE TABLE IF NOT EXISTS playback_history (
                                                id SERIAL PRIMARY KEY,
                                                user_id INT NOT NULL REFERENCES users(id),
    song_id INT NOT NULL REFERENCES songs(id),
    played_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS playlists (
                                         id SERIAL PRIMARY KEY,
                                         user_id INT NOT NULL REFERENCES users(id),
    song_id INT NOT NULL REFERENCES songs(id),
    position INT NOT NULL
    );

CREATE TABLE IF NOT EXISTS stations (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(100) NOT NULL,
    tags TEXT NOT NULL -- storing tags as a comma-separated string
    );

CREATE TABLE plays (
                       id SERIAL PRIMARY KEY,
                       user_id INT REFERENCES users(id),
                       song_id INT REFERENCES songs(id),
                       action VARCHAR(50) NOT NULL, -- e.g., start, pause, stop
                       timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE play_aggregates (
                                 song_id INT REFERENCES songs(id),
                                 user_id INT REFERENCES users(id),
                                 total_plays INT DEFAULT 0,
                                 total_duration INT DEFAULT 0, -- duration in seconds
                                 PRIMARY KEY (song_id, user_id)
);

CREATE TABLE song_play_aggregates (
                                      song_id INT REFERENCES songs(id),
                                      total_plays INT DEFAULT 0,
                                      total_duration INT DEFAULT 0, -- duration in seconds
                                      PRIMARY KEY (song_id)
);
