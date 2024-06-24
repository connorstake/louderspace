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
    ('Synth Beats 1', 'Artist 4', ARRAY['synth', 'instrumental', 'beats', 'lofi', 'hiphop'], '6e3cd1cf-f487-487f-b53b-e858b9b101eb', TRUE),
    ('Instrumental Vibes', 'Artist 5', ARRAY['synth', 'instrumental', 'beats', 'lofi', 'hiphop'], 'eee20928-b7ab-44f5-8a9f-31f2f7d568a5', TRUE),
    ('Classical Hiphop', 'Artist 6', ARRAY['synth', 'instrumental', 'beats', 'lofi', 'hiphop'], '6d436627-10c0-4a02-a384-5d13a27d8d96', TRUE),
    ('Piano Lo-fi', 'Artist 7', ARRAY['synth', 'instrumental', 'beats', 'lofi', 'hiphop'], '3c6534d5-fab4-4e9a-b230-471a76debcbf', TRUE),
    ('Synthwave Chill', 'Artist 8', ARRAY['synth', 'instrumental', 'beats', 'lofi', 'hiphop'], '7be45898-26a1-479b-bbbd-aaa39ec83551', TRUE),
    ('Instrumental Beats', 'Artist 9', ARRAY['synth', 'instrumental', 'beats', 'lofi', 'hiphop'], 'b089275f-f73f-4094-b5ae-77e98b1c3311', TRUE),
    ('Song 1', 'Random Artist', ARRAY['classical', 'lofi'], '6033d8c5-e024-4d84-80a1-1df28683b304', TRUE),
    ('Song 2', 'Random Artist', ARRAY['classical', 'lofi'], 'c7b04693-164a-448e-a98b-1ce02bf750a1', TRUE),
    ('Song 3', 'Random Artist', ARRAY['classical', 'lofi'], '6f3d4e9f-90a7-4e87-b6bd-5d0b67085b25', TRUE),
    ('Song 4', 'Random Artist', ARRAY['classical', 'lofi'], 'e37d4828-44f4-4394-93fe-fcb4eee6619e', TRUE),
    ('Song 5', 'Random Artist', ARRAY['classical', 'lofi'], 'd8c2a782-5d34-455d-8f67-f030172a3e69', TRUE),
    ('Song 6', 'Random Artist', ARRAY['classical', 'lofi'], '67a74714-bc38-4247-b936-a2bf29f7854f', TRUE),
    ('Song 7', 'Random Artist', ARRAY['classical', 'lofi'], '96edec8b-97af-47ab-a1ee-0cb692725d4f', TRUE),
    ('Song 8', 'Random Artist', ARRAY['classical', 'lofi'], 'b0959337-1f4c-447d-88b6-32bf50bbfa90', TRUE),
    ('Song 9', 'Random Artist', ARRAY['classical', 'lofi'], '6e4b0a38-cbe3-469f-98d4-74ef9099a2b2', TRUE),
    ('Song 10', 'Random Artist', ARRAY['classical', 'lofi'], '3073ec18-bb02-4a15-b8a8-223680e83bd8', TRUE),
    ('Song 11', 'Random Artist', ARRAY['classical', 'lofi'], '47b44acb-9432-4aef-b26d-1e1b665729a1', TRUE),
    ('Song 12', 'Random Artist', ARRAY['classical', 'lofi'], 'adc3ca43-cb55-41c0-ae7f-95a002690939', TRUE),
    ('Song 13', 'Random Artist', ARRAY['classical', 'lofi'], '29a278fc-1ad5-4bdb-98a3-50cff1c60ae4', TRUE),
    ('Song 14', 'Random Artist', ARRAY['classical', 'lofi'], 'ed4bfcc3-f68b-41e1-8199-73017b7c44f3', TRUE),
    ('Song 15', 'Random Artist', ARRAY['classical', 'lofi'], '1d5128b8-ee1d-4b9b-804a-09ef79da5ece', TRUE),
    ('Song 16', 'Random Artist', ARRAY['classical', 'lofi'], 'ae4094c1-36aa-489b-a046-f8b2ef9da914', TRUE);

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
    (6, 5), -- Instrumental Beats -> hiphop
    (7, 8), -- Song 1 -> classical
    (7, 4), -- Song 1 -> lofi
    (8, 8), -- Song 2 -> classical
    (8, 4), -- Song 2 -> lofi
    (9, 8), -- Song 3 -> classical
    (9, 4), -- Song 3 -> lofi
    (10, 8), -- Song 4 -> classical
    (10, 4), -- Song 4 -> lofi
    (11, 8), -- Song 5 -> classical
    (11, 4), -- Song 5 -> lofi
    (12, 8), -- Song 6 -> classical
    (12, 4), -- Song 6 -> lofi
    (13, 8), -- Song 7 -> classical
    (13, 4), -- Song 7 -> lofi
    (14, 8), -- Song 8 -> classical
    (14, 4), -- Song 8 -> lofi
    (15, 8), -- Song 9 -> classical
    (15, 4), -- Song 9 -> lofi
    (16, 8), -- Song 10 -> classical
    (16, 4), -- Song 10 -> lofi
    (17, 8), -- Song 11 -> classical
    (17, 4), -- Song 11 -> lofi
    (18, 8), -- Song 12 -> classical
    (18, 4), -- Song 12 -> lofi
    (19, 8), -- Song 13 -> classical
    (19, 4), -- Song 13 -> lofi
    (20, 8), -- Song 14 -> classical
    (20, 4), -- Song 14 -> lofi
    (21, 8), -- Song 15 -> classical
    (21, 4), -- Song 15 -> lofi
    (22, 8), -- Song 16 -> classical
    (22, 4); -- Song 16 -> lofi

-- Insert initial stations
INSERT INTO stations (name, tags)
VALUES
    ('Synthy Lo-fi', 'synth, lofi'),
    ('Classical Lo-fi', 'classical, lofi');
