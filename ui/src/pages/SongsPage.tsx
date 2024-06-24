import React, { useEffect, useState } from 'react';
import { useParams, useLocation } from 'react-router-dom';
import { Container, Typography, Grid, Card, CardContent, CardActions, Button } from '@mui/material';
import axios from 'axios';

interface Song {
    id: number;
    title: string;
    artist: string;
    genre: string;
    sunoApiId: string;
}

const SongsPage: React.FC = () => {
    const [songs, setSongs] = useState<Song[]>([]);
    const [currentSong, setCurrentSong] = useState<Song | null>(null);
    const { stationId } = useParams<{ stationId?: string }>();
    const location = useLocation();

    useEffect(() => {
        const fetchSongs = async () => {
            try {
                let response;
                if (stationId) {
                    response = await axios.get<Song[]>(`http://localhost:8080/stations/${stationId}/songs`);
                } else {
                    response = await axios.get<Song[]>('http://localhost:8080/songs');
                }
                setSongs(response.data);
            } catch (error) {
                console.error('There was an error fetching the songs!', error);
            }
        };

        fetchSongs();
    }, [stationId]);

    return (
        <Container>
            <Typography variant="h4" component="h1" gutterBottom>
                {stationId ? 'Songs for Station' : 'All Songs'}
            </Typography>
            <Grid container spacing={4}>
                {songs.map((song) => (
                    <Grid item key={song.id} xs={12} sm={6} md={4}>
                        <Card>
                            <CardContent>
                                <Typography variant="h5" component="div">
                                    {song.title}
                                </Typography>
                                <Typography variant="body2" color="text.secondary">
                                    Artist: {song.artist}
                                </Typography>
                                <Typography variant="body2" color="text.secondary">
                                    Genre: {song.genre}
                                </Typography>
                            </CardContent>
                            <CardActions>
                                <Button size="small" onClick={() => handlePlaySong(song)}>Play</Button>
                            </CardActions>
                        </Card>
                    </Grid>
                ))}
            </Grid>
            {currentSong && (
                <div style={{ marginTop: '20px' }}>
                    <Typography variant="h5" component="div">
                        Now Playing: {currentSong.title} by {currentSong.artist}
                    </Typography>
                    <audio controls autoPlay>
                        <source src={`https://cdn.sunoapi.com/${currentSong.sunoApiId}.mp3`} type="audio/mpeg" />
                        Your browser does not support the audio element.
                    </audio>
                </div>
            )}
        </Container>
    );

    function handlePlaySong(song: Song) {
        setCurrentSong(song);
    }
};

export default SongsPage;
