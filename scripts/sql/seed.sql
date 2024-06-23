-- Insert initial users
INSERT INTO users (username, password, email)
VALUES
    ('user1', 'password1', 'user1@example.com'),
    ('user2', 'password2', 'user2@example.com');

-- Insert initial tags
INSERT INTO tags (name)
VALUES
    ('chill'),
    ('beats'),
    ('vibes'),
    ('lofi'),
    ('hiphop');

-- Insert initial songs
INSERT INTO songs (title, artist, genre, suno_id, is_generated)
VALUES
    ('Chill Song 1', 'Artist 1', 'chill, beats', 'suno1', TRUE),
    ('Chill Song 2', 'Artist 2', 'chill, vibes', 'suno2', TRUE),
    ('Lo-fi Song 1', 'Artist 3', 'lofi, hiphop', 'suno3', TRUE);

-- Insert song tags relationships
INSERT INTO song_tags (song_id, tag_id)
VALUES
    (1, 1), -- Chill Song 1 -> chill
    (1, 2), -- Chill Song 1 -> beats
    (2, 1), -- Chill Song 2 -> chill
    (2, 3), -- Chill Song 2 -> vibes
    (3, 4), -- Lo-fi Song 1 -> lofi
    (3, 5); -- Lo-fi Song 1 -> hiphop

-- Insert initial stations
INSERT INTO stations (name, tags)
VALUES
    ('Chill Station', 'chill, beats'),
    ('Lo-fi Station', 'lofi, hiphop');
