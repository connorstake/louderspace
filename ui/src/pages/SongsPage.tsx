import React, { useState, useRef } from 'react';
import { Container, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, TextField, Select, MenuItem, FormControl, InputLabel, OutlinedInput, Box, Chip } from '@mui/material';
import useSongs from '../hooks/useSongs';

interface Song {
    id: number;
    title: string;
    artist: string;
    genre: string;
    suno_id: string;
    tags: Tag[];  // Add tags to the Song interface
}

interface Tag {
    id: number;
    name: string;
}

const SongsPage: React.FC = () => {
    const { songs, tags: availableTags, loading, error, addSong, updateExistingSong, removeSong } = useSongs();
    const [currentSong, setCurrentSong] = useState<Song | null>(null);
    const [openAdd, setOpenAdd] = useState(false);
    const [openEdit, setOpenEdit] = useState(false);
    const [openDelete, setOpenDelete] = useState(false);
    const [newSong, setNewSong] = useState<Song>({ id: 0, title: '', artist: '', genre: '', suno_id: '', tags: [] });
    const [selectedSong, setSelectedSong] = useState<Song | null>(null);
    const audioRef = useRef<HTMLAudioElement>(null);
    const [newSongTags, setNewSongTags] = useState<string[]>([]);

    const handlePlaySong = (song: Song) => {
        console.log('Playing song:', song);
        setCurrentSong(song);

        // Reload the audio element to play the new song
        if (audioRef.current) {
            audioRef.current.load();
        }
    };

    const handleOpenAdd = () => {
        setNewSong({ id: 0, title: '', artist: '', genre: '', suno_id: '', tags: [] });
        setOpenAdd(true);
    };

    const handleOpenEdit = (song: Song) => {
        setSelectedSong(song);
        setNewSong(song);
        setNewSongTags(song.tags ? song.tags.map(tag => tag.name) : []);
        setOpenEdit(true);
    };

    const handleOpenDelete = (song: Song) => {
        setSelectedSong(song);
        setOpenDelete(true);
    };

    const handleCloseAdd = () => {
        setOpenAdd(false);
        setNewSongTags([]);
    };

    const handleCloseEdit = () => {
        setOpenEdit(false);
    };

    const handleCloseDelete = () => {
        setOpenDelete(false);
    };

    const handleAddSong = async () => {
        await addSong({ ...newSong, tags: newSongTags });
        setNewSongTags([]);
        handleCloseAdd();
    };

    const handleEditSong = async () => {
        if (selectedSong !== null) {
            await updateExistingSong(selectedSong.id, { ...newSong, tags: newSongTags });
            setNewSongTags([]);
            handleCloseEdit();
        }
    };

    const handleDeleteSong = async () => {
        if (selectedSong !== null) {
            await removeSong(selectedSong.id);
            handleCloseDelete();
        }
    };

    return (
        <Container>
            <Typography variant="h4" component="h1" gutterBottom>
                Songs
            </Typography>
            <Button variant="contained" color="primary" onClick={handleOpenAdd} sx={{ mb: 4 }}>
                Add Song
            </Button>
            {loading ? (
                <Typography>Loading...</Typography>
            ) : error ? (
                <Typography color="error">{error}</Typography>
            ) : (
                <TableContainer component={Paper}>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell>Title</TableCell>
                                <TableCell>Artist</TableCell>
                                <TableCell>Genre</TableCell>
                                <TableCell>Tags</TableCell>
                                <TableCell>Action</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {songs.map((song) => (
                                <TableRow key={song.id}>
                                    <TableCell>{song.title}</TableCell>
                                    <TableCell>{song.artist}</TableCell>
                                    <TableCell>{song.genre}</TableCell>
                                    <TableCell>{song.tags && song.tags.map((tag)=>tag.name).join(',')}</TableCell>
                                    <TableCell>
                                        <Button size="small" onClick={() => handlePlaySong(song)}>Play</Button>
                                        <Button size="small" color="primary" onClick={() => handleOpenEdit(song)}>Edit</Button>
                                        <Button size="small" color="secondary" onClick={() => handleOpenDelete(song)}>Delete</Button>
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            )}
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

            <Dialog open={openAdd} onClose={handleCloseAdd}>
                <DialogTitle>Add Song</DialogTitle>
                <DialogContent>
                    <TextField
                        label="Title"
                        value={newSong.title}
                        onChange={(e) => setNewSong({ ...newSong, title: e.target.value })}
                        fullWidth
                        margin="normal"
                    />
                    <TextField
                        label="Artist"
                        value={newSong.artist}
                        onChange={(e) => setNewSong({ ...newSong, artist: e.target.value })}
                        fullWidth
                        margin="normal"
                    />
                    <TextField
                        label="Genre"
                        value={newSong.genre}
                        onChange={(e) => setNewSong({ ...newSong, genre: e.target.value })}
                        fullWidth
                        margin="normal"
                    />
                    <TextField
                        label="Suno ID"
                        value={newSong.suno_id}
                        onChange={(e) => setNewSong({ ...newSong, suno_id: e.target.value })}
                        fullWidth
                        margin="normal"
                    />
                    <FormControl fullWidth margin="normal">
                        <InputLabel id="tags-label">Tags</InputLabel>
                        <Select
                            labelId="tags-label"
                            multiple
                            value={newSongTags}
                            onChange={(e) => setNewSongTags(e.target.value as string[])}
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
                    <Button onClick={handleAddSong} color="primary">
                        Add
                    </Button>
                </DialogActions>
            </Dialog>

            <Dialog open={openEdit} onClose={handleCloseEdit}>
                <DialogTitle>Edit Song</DialogTitle>
                <DialogContent>
                    <TextField
                        label="Title"
                        value={newSong.title}
                        onChange={(e) => setNewSong({ ...newSong, title: e.target.value })}
                        fullWidth
                        margin="normal"
                    />
                    <TextField
                        label="Artist"
                        value={newSong.artist}
                        onChange={(e) => setNewSong({ ...newSong, artist: e.target.value })}
                        fullWidth
                        margin="normal"
                    />
                    <TextField
                        label="Genre"
                        value={newSong.genre}
                        onChange={(e) => setNewSong({ ...newSong, genre: e.target.value })}
                        fullWidth
                        margin="normal"
                    />
                    <TextField
                        label="Suno ID"
                        value={newSong.suno_id}
                        onChange={(e) => setNewSong({ ...newSong, suno_id: e.target.value })}
                        fullWidth
                        margin="normal"
                    />
                    <FormControl fullWidth margin="normal">
                        <InputLabel id="tags-label">Tags</InputLabel>
                        <Select
                            labelId="tags-label"
                            multiple
                            value={newSongTags}
                            onChange={(e) => setNewSongTags(e.target.value as string[])}
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
                    <Button onClick={handleCloseEdit} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={handleEditSong} color="primary">
                        Save
                    </Button>
                </DialogActions>
            </Dialog>

            <Dialog open={openDelete} onClose={handleCloseDelete}>
                <DialogTitle>Delete Song</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        Are you sure you want to delete this song? This action cannot be undone.
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseDelete} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={handleDeleteSong} color="secondary">
                        Delete
                    </Button>
                </DialogActions>
            </Dialog>
        </Container>
    );
};

export default SongsPage;
