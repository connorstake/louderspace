import React, { useEffect, useState, useRef } from 'react';
import { useParams } from 'react-router-dom';
import { Container, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, Button } from '@mui/material';
import axios from 'axios';

interface Song {
    id: number;
    title: string;
    artist: string;
    genre: string;
    suno_id: string;
}

const SongsPage: React.FC = () => {
    const [songs, setSongs] = useState<Song[]>([]);
    const [currentSong, setCurrentSong] = useState<Song | null>(null);
    const audioRef = useRef<HTMLAudioElement>(null);
    const { stationId } = useParams<{ stationId?: string }>();

    useEffect(() => {
        const fetchSongs = async () => {
            try {
                let response;
                if (stationId) {
                    response = await axios.get<Song[]>(`http://localhost:8080/stations/${stationId}/songs`);
                } else {
                    response = await axios.get<Song[]>('http://localhost:8080/songs');
                }
                console.log('Fetched songs:', response.data);
                setSongs(response.data);
            } catch (error) {
                console.error('There was an error fetching the songs!', error);
            }
        };

        fetchSongs();
    }, [stationId]);

    const handlePlaySong = (song: Song) => {
        console.log('Playing song:', song);
        setCurrentSong(song);

        // Reload the audio element to play the new song
        if (audioRef.current) {
            audioRef.current.load();
        }
    };

    return (
        <Container>
            <Typography variant="h4" component="h1" gutterBottom>
                {stationId ? 'Songs for Station' : 'All Songs'}
            </Typography>
            <TableContainer component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Title</TableCell>
                            <TableCell>Artist</TableCell>
                            <TableCell>Genre</TableCell>
                            <TableCell>Action</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {songs.map((song) => (
                            <TableRow key={song.id}>
                                <TableCell>{song.title}</TableCell>
                                <TableCell>{song.artist}</TableCell>
                                <TableCell>{song.genre}</TableCell>
                                <TableCell>
                                    <Button size="small" onClick={() => handlePlaySong(song)}>Play</Button>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
            {currentSong && (
                <div style={{ marginTop: '20px' }}>
                    <Typography variant="h5" component="div">
                        Now Playing: {currentSong.title} by {currentSong.artist}
                    </Typography>
                    <audio controls autoPlay ref={audioRef} onError={() => console.error('Error loading audio')}>
                        <source src={`https://cdn1.suno.ai/${currentSong.suno_id}.mp3?api_key=${process.env.REACT_APP_API_KEY}`} type="audio/mpeg" />
                        Your browser does not support the audio element.
                    </audio>
                </div>
            )}
        </Container>
    );
};

export default SongsPage;
