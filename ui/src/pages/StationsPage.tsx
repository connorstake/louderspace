import React, { useEffect, useState } from 'react';
import { Container, Typography, Grid, Card, CardContent, CardActions, Button, TextField, Box, Select, MenuItem, FormControl, InputLabel, OutlinedInput, Chip } from '@mui/material';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

interface Station {
    id: number;
    name: string;
    tags: string[];
}

const StationsPage: React.FC = () => {
    const [stations, setStations] = useState<Station[]>([]);
    const [newStationName, setNewStationName] = useState<string>('');
    const [newStationTags, setNewStationTags] = useState<string[]>([]);
    const [availableTags, setAvailableTags] = useState<string[]>([]);
    const navigate = useNavigate();

    useEffect(() => {
        // Fetch stations from the backend
        axios.get<Station[]>('http://localhost:8080/stations')
            .then(response => {
                setStations(response.data);
            })
            .catch(error => {
                console.error('There was an error fetching the stations!', error);
            });

        // Fetch available tags from the backend
        axios.get<string[]>('http://localhost:8080/tags')
            .then(response => {
                setAvailableTags(response.data);
            })
            .catch(error => {
                console.error('There was an error fetching the tags!', error);
            });
    }, []);

    const handleAddStation = () => {
        axios.post('http://localhost:8080/stations', { name: newStationName, tags: newStationTags })
            .then(response => {
                setStations(prevStations => [...prevStations, response.data]);
                setNewStationName('');
                setNewStationTags([]);
            })
            .catch(error => {
                console.error('There was an error adding the station!', error);
            });
    };

    return (
        <Container>
            <Typography variant="h4" component="h1" gutterBottom>
                Stations
            </Typography>
            <Box component="form" sx={{ mb: 4 }}>
                <TextField
                    label="Station Name"
                    value={newStationName}
                    onChange={(e) => setNewStationName(e.target.value)}
                    fullWidth
                    margin="normal"
                />
                <FormControl fullWidth margin="normal">
                    <InputLabel id="tags-label">Tags</InputLabel>
                    <Select
                        labelId="tags-label"
                        multiple
                        value={newStationTags}
                        onChange={(e) => setNewStationTags(e.target.value as string[])}
                        input={<OutlinedInput id="select-multiple-chip" label="Tags" />}
                        renderValue={(selected) => (
                            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                                {selected.map((value) => (
                                    <Chip key={value} label={value} />
                                ))}
                            </Box>
                        )}
                    >
                        {availableTags.map((tag) => (
                            <MenuItem key={tag} value={tag}>
                                {tag}
                            </MenuItem>
                        ))}
                    </Select>
                </FormControl>
                <Button variant="contained" color="primary" onClick={handleAddStation} sx={{ mt: 2 }}>
                    Add Station
                </Button>
            </Box>
            <Grid container spacing={4}>
                {stations.map((station) => (
                    <Grid item key={station.id} xs={12} sm={6} md={4}>
                        <Card>
                            <CardContent>
                                <Typography variant="h5" component="div">
                                    {station.name}
                                </Typography>
                                <Typography variant="body2" color="text.secondary">
                                    Tags: {station.tags ? station.tags.join(', ') : 'No tags'}
                                </Typography>
                            </CardContent>
                            <CardActions>
                                <Button size="small" onClick={() => navigate(`/stations/${station.id}/songs`)}>View Songs</Button>
                            </CardActions>
                        </Card>
                    </Grid>
                ))}
            </Grid>
        </Container>
    );
};

export default StationsPage;
