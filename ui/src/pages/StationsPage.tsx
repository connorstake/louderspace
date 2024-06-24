import React, { useEffect, useState } from 'react';
import { Container, Typography, Grid, Card, CardContent, CardActions, Button, TextField, Box, Select, MenuItem, FormControl, InputLabel, OutlinedInput, Chip, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle } from '@mui/material';
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
    const [selectedStationId, setSelectedStationId] = useState<number | null>(null);
    const [openDelete, setOpenDelete] = useState(false);
    const [openAdd, setOpenAdd] = useState(false);
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
                setOpenAdd(false);
            })
            .catch(error => {
                console.error('There was an error adding the station!', error);
            });
    };

    const handleDeleteStation = () => {
        if (selectedStationId !== null) {
            axios.delete(`http://localhost:8080/stations/${selectedStationId}`)
                .then(() => {
                    setStations(prevStations => prevStations.filter(station => station.id !== selectedStationId));
                    setSelectedStationId(null);
                    setOpenDelete(false);
                })
                .catch(error => {
                    console.error('There was an error deleting the station!', error);
                });
        }
    };

    const handleClickOpenDelete = (stationId: number) => {
        setSelectedStationId(stationId);
        setOpenDelete(true);
    };

    const handleClickOpenAdd = () => {
        setOpenAdd(true);
    };

    const handleCloseDelete = () => {
        setSelectedStationId(null);
        setOpenDelete(false);
    };

    const handleCloseAdd = () => {
        setNewStationName('');
        setNewStationTags([]);
        setOpenAdd(false);
    };

    return (
        <Container>
            <Typography variant="h4" component="h1" gutterBottom>
                Stations
            </Typography>
            <Button variant="contained" color="primary" onClick={handleClickOpenAdd} sx={{ mb: 4 }}>
                Add Station
            </Button>
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
                                <Button size="small" color="secondary" onClick={() => handleClickOpenDelete(station.id)}>Delete</Button>
                            </CardActions>
                        </Card>
                    </Grid>
                ))}
            </Grid>
            <Dialog open={openDelete} onClose={handleCloseDelete}>
                <DialogTitle>Delete Station</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        Are you sure you want to delete this station? This action cannot be undone.
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseDelete} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={handleDeleteStation} color="secondary">
                        Delete
                    </Button>
                </DialogActions>
            </Dialog>
            <Dialog open={openAdd} onClose={handleCloseAdd}>
                <DialogTitle>Add Station</DialogTitle>
                <DialogContent>
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
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseAdd} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={handleAddStation} color="primary">
                        Add
                    </Button>
                </DialogActions>
            </Dialog>
        </Container>
    );
};

export default StationsPage;
