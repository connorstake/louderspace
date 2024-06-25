import React, { useState } from 'react';
import {
    Container, Typography, Button, TextField, Dialog, DialogActions, DialogContent,
    DialogContentText, DialogTitle, Table, TableBody, TableCell, TableContainer, TableHead,
    TableRow, Paper, IconButton
} from '@mui/material';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import useTags, { Tag } from '../hooks/useTags';

const TagsPage: React.FC = () => {
    const { tags, loading, error, addTag, updateExistingTag, removeTag } = useTags();
    const [newTagName, setNewTagName] = useState<string>('');
    const [selectedTag, setSelectedTag] = useState<Tag | null>(null);
    const [openAdd, setOpenAdd] = useState(false);
    const [openEdit, setOpenEdit] = useState(false);
    const [openDelete, setOpenDelete] = useState(false);

    const handleAddTag = async () => {
        await addTag({ name: newTagName });
        setNewTagName('');
        setOpenAdd(false);
    };

    const handleEditTag = async () => {
        if (selectedTag) {
            await updateExistingTag(selectedTag.id, { name: newTagName });
            setSelectedTag(null);
            setNewTagName('');
            setOpenEdit(false);
        }
    };

    const handleDeleteTag = async () => {
        if (selectedTag) {
            await removeTag(selectedTag.id);
            setSelectedTag(null);
            setOpenDelete(false);
        }
    };

    const handleClickOpenAdd = () => {
        setOpenAdd(true);
    };

    const handleClickOpenEdit = (tag: Tag) => {
        setSelectedTag(tag);
        setNewTagName(tag.name);
        setOpenEdit(true);
    };

    const handleClickOpenDelete = (tag: Tag) => {
        setSelectedTag(tag);
        setOpenDelete(true);
    };

    const handleCloseAdd = () => {
        setNewTagName('');
        setOpenAdd(false);
    };

    const handleCloseEdit = () => {
        setSelectedTag(null);
        setNewTagName('');
        setOpenEdit(false);
    };

    const handleCloseDelete = () => {
        setSelectedTag(null);
        setOpenDelete(false);
    };

    return (
        <Container>
            <Typography variant="h4" component="h1" gutterBottom>
                Tags
            </Typography>
            <Button variant="contained" color="primary" onClick={handleClickOpenAdd} sx={{ mb: 4 }}>
                Add Tag
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
                                <TableCell>Name</TableCell>
                                <TableCell align="right">Actions</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {tags.map((tag) => (
                                <TableRow key={tag.id}>
                                    <TableCell>{tag.name}</TableCell>
                                    <TableCell align="right">
                                        <IconButton color="primary" onClick={() => handleClickOpenEdit(tag)}>
                                            <EditIcon />
                                        </IconButton>
                                        <IconButton color="secondary" onClick={() => handleClickOpenDelete(tag)}>
                                            <DeleteIcon />
                                        </IconButton>
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            )}
            <Dialog open={openAdd} onClose={handleCloseAdd}>
                <DialogTitle>Add Tag</DialogTitle>
                <DialogContent>
                    <TextField
                        label="Tag Name"
                        value={newTagName}
                        onChange={(e) => setNewTagName(e.target.value)}
                        fullWidth
                        margin="normal"
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseAdd} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={handleAddTag} color="primary">
                        Add
                    </Button>
                </DialogActions>
            </Dialog>
            <Dialog open={openEdit} onClose={handleCloseEdit}>
                <DialogTitle>Edit Tag</DialogTitle>
                <DialogContent>
                    <TextField
                        label="Tag Name"
                        value={newTagName}
                        onChange={(e) => setNewTagName(e.target.value)}
                        fullWidth
                        margin="normal"
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseEdit} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={handleEditTag} color="primary">
                        Save
                    </Button>
                </DialogActions>
            </Dialog>
            <Dialog open={openDelete} onClose={handleCloseDelete}>
                <DialogTitle>Delete Tag</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        Are you sure you want to delete this tag? This action cannot be undone.
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseDelete} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={handleDeleteTag} color="secondary">
                        Delete
                    </Button>
                </DialogActions>
            </Dialog>
        </Container>
    );
};

export default TagsPage;
