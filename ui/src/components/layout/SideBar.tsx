// src/components/layout/Sidebar.tsx
import React from 'react';
import { Drawer, List, ListItem, ListItemText } from '@mui/material';
import { Link } from 'react-router-dom';

const Sidebar: React.FC = () => {
    return (
        <Drawer variant="permanent" anchor="left">
            <List>
                <ListItem button component={Link} to="/">
                    <ListItemText primary="Dashboard" />
                </ListItem>
                <ListItem button component={Link} to="/users">
                    <ListItemText primary="Users" />
                </ListItem>
                <ListItem button component={Link} to="/stations">
                    <ListItemText primary="Stations" />
                </ListItem>
                <ListItem button component={Link} to="/songs">
                    <ListItemText primary="Songs" />
                </ListItem>
                <ListItem button component={Link} to="/settings">
                    <ListItemText primary="Settings" />
                </ListItem>
            </List>
        </Drawer>
    );
};

export default Sidebar;
