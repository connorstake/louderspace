-- Insert initial users
INSERT INTO users (username, password, email)
VALUES
    ('user1', 'password1', 'user1@example.com'),
    ('user2', 'password2', 'user2@example.com');

-- Insert new tags
INSERT INTO tags (name)
VALUES
    ('chill'),
    ('beats'),
    ('vibes'),
    ('lofi'),
    ('hiphop'),
    ('synth'),
    ('instrumental'),
    ('classical'),
    ('piano');

-- Insert new songs
INSERT INTO songs (title, artist, genre, suno_id, is_generated)
VALUES
    ('Synth Beats 1', 'Artist 4', 'synth, instrumental, beats, lofi, hiphop', '6e3cd1cf-f487-487f-b53b-e858b9b101eb', TRUE),
    ('Instrumental Vibes', 'Artist 5', 'synth, instrumental, beats, lofi, hiphop', 'eee20928-b7ab-44f5-8a9f-31f2f7d568a5', TRUE),
    ('Classical Hiphop', 'Artist 6', 'synth, instrumental, beats, lofi, hiphop', '6d436627-10c0-4a02-a384-5d13a27d8d96', TRUE),
    ('Piano Lo-fi', 'Artist 7', 'synth, instrumental, beats, lofi, hiphop', '3c6534d5-fab4-4e9a-b230-471a76debcbf', TRUE),
    ('Synthwave Chill', 'Artist 8', 'synth, instrumental, beats, lofi, hiphop', '7be45898-26a1-479b-bbbd-aaa39ec83551', TRUE),
    ('Instrumental Beats', 'Artist 9', 'synth, instrumental, beats, lofi, hiphop', 'b089275f-f73f-4094-b5ae-77e98b1c3311', TRUE);

-- Insert song tags relationships for new songs
-- Assuming tag ids for the new tags are 6, 7, 8, 9 and existing ones for beats (2), lofi (4), hiphop (5)
INSERT INTO song_tags (song_id, tag_id)
VALUES
    (1, 6), -- Synth Beats 1 -> synth
    (1, 7), -- Synth Beats 1 -> instrumental
    (1, 2), -- Synth Beats 1 -> beats
    (1, 4), -- Synth Beats 1 -> lofi
    (1, 5), -- Synth Beats 1 -> hiphop
    (2, 6), -- Instrumental Vibes -> synth
    (2, 7), -- Instrumental Vibes -> instrumental
    (2, 2), -- Instrumental Vibes -> beats
    (2, 4), -- Instrumental Vibes -> lofi
    (2, 5), -- Instrumental Vibes -> hiphop
    (3, 6), -- Classical Hiphop -> synth
    (3, 7), -- Classical Hiphop -> instrumental
    (3, 2), -- Classical Hiphop -> beats
    (3, 4), -- Classical Hiphop -> lofi
    (3, 5), -- Classical Hiphop -> hiphop
    (4, 6), -- Piano Lo-fi -> synth
    (4, 7), -- Piano Lo-fi -> instrumental
    (4, 2), -- Piano Lo-fi -> beats
    (4, 4), -- Piano Lo-fi -> lofi
    (4, 5), -- Piano Lo-fi -> hiphop
    (5, 6), -- Synthwave Chill -> synth
    (5, 7), -- Synthwave Chill -> instrumental
    (5, 2), -- Synthwave Chill -> beats
    (5, 4), -- Synthwave Chill -> lofi
    (5, 5), -- Synthwave Chill -> hiphop
    (6, 6), -- Instrumental Beats -> synth
    (6, 7), -- Instrumental Beats -> instrumental
    (6, 2), -- Instrumental Beats -> beats
    (6, 4), -- Instrumental Beats -> lofi
    (6, 5); -- Instrumental Beats -> hiphop

-- Insert initial stations
INSERT INTO stations (name, tags)
VALUES
    ('Chill Station', 'chill, beats'),
    ('Lo-fi Station', 'lofi, hiphop');
