import React, { useState } from 'react';
import {
    Container, Typography, Grid, Card, CardContent, CardActions, Button, TextField, Box, Select, MenuItem, FormControl,
    InputLabel, OutlinedInput, Chip, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, IconButton
} from '@mui/material';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import { useNavigate } from 'react-router-dom';
import useStations from '../hooks/useStation';

interface Station {
    id: number;
    name: string;
    tags: string[];
}

const StationsPage: React.FC = () => {
    const {
        stations,
        tags: availableTags,
        loading,
        error,
        addStation,
        updateExistingStation,
        removeStation,
    } = useStations();

    const [newStationName, setNewStationName] = useState<string>('');
    const [newStationTags, setNewStationTags] = useState<string[]>([]);
    const [selectedStation, setSelectedStation] = useState<Station | null>(null);
    const [openDelete, setOpenDelete] = useState(false);
    const [openAdd, setOpenAdd] = useState(false);
    const [openEdit, setOpenEdit] = useState(false);
    const navigate = useNavigate();

    const handleAddStation = () => {
        addStation(newStationName, newStationTags);
        setNewStationName('');
        setNewStationTags([]);
        setOpenAdd(false);
    };

    const handleUpdateStation = () => {
        if (selectedStation !== null) {
            updateExistingStation(selectedStation.id, newStationName, newStationTags);
            setSelectedStation(null);
            setNewStationName('');
            setNewStationTags([]);
            setOpenEdit(false);
        }
    };

    const handleDeleteStation = () => {
        if (selectedStation !== null) {
            removeStation(selectedStation.id);
            setSelectedStation(null);
            setOpenDelete(false);
        }
    };

    const handleClickOpenDelete = (station: Station) => {
        setSelectedStation(station);
        setOpenDelete(true);
    };

    const handleClickOpenAdd = () => {
        setOpenAdd(true);
    };

    const handleClickOpenEdit = (station: Station) => {
        setSelectedStation(station);
        setNewStationName(station.name);
        setNewStationTags(station.tags);
        setOpenEdit(true);
    };

    const handleCloseDelete = () => {
        setSelectedStation(null);
        setOpenDelete(false);
    };

    const handleCloseAdd = () => {
        setNewStationName('');
        setNewStationTags([]);
        setOpenAdd(false);
    };

    const handleCloseEdit = () => {
        setSelectedStation(null);
        setNewStationName('');
        setNewStationTags([]);
        setOpenEdit(false);
    };

    return (
        <Container>
            <Typography variant="h4" component="h1" gutterBottom>
                Stations
            </Typography>
            <Button variant="contained" color="primary" onClick={handleClickOpenAdd} sx={{ mb: 4 }}>
                Add Station
            </Button>
            {loading ? (
                <Typography>Loading...</Typography>
            ) : error ? (
                <Typography color="error">{error}</Typography>
            ) : (
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
                                    <IconButton color="primary" onClick={() => handleClickOpenEdit(station)}>
                                        <EditIcon />
                                    </IconButton>
                                    <IconButton color="secondary" onClick={() => handleClickOpenDelete(station)}>
                                        <DeleteIcon />
                                    </IconButton>
                                </CardActions>
                            </Card>
                        </Grid>
                    ))}
                </Grid>
            )}
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
                                <MenuItem key={tag.id} value={tag.name}>
                                    {tag.name}
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
            <Dialog open={openEdit} onClose={handleCloseEdit}>
                <DialogTitle>Edit Station</DialogTitle>
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
                                        <Chip key={value} label={value} onDelete={() => setNewStationTags(prevTags => prevTags.filter(tag => tag !== value))}/>
                                    ))}
                                </Box>
                            )}
                        >
                            {availableTags.map((tag) => (
                                <MenuItem key={tag.id} value={tag.name}>
                                    {tag.name}
                                </MenuItem>
                            ))}
                        </Select>
                    </FormControl>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseEdit} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={handleUpdateStation} color="primary">
                        Save
                    </Button>
                </DialogActions>
            </Dialog>
        </Container>
    );
};

export default StationsPage;
