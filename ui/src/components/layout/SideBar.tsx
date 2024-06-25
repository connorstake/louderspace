import React from 'react';
import { Link } from 'react-router-dom';
import { List, ListItem, ListItemText } from '@mui/material';

const SideBar: React.FC = () => {
    return (
        <div style={{ width: '250px', height: '100vh', backgroundColor: '#f0f0f0', padding: '1rem' }}>
            <List>
                <ListItem component={Link} to="/">
                    <ListItemText primary="Dashboard" />
                </ListItem>
                <ListItem  component={Link} to="/users">
                    <ListItemText primary="Users" />
                </ListItem>
                <ListItem  component={Link} to="/stations">
                    <ListItemText primary="Stations" />
                </ListItem>
                <ListItem  component={Link} to="/songs">
                    <ListItemText primary="Songs" />
                </ListItem>
                <ListItem  component={Link} to="/tags">
                    <ListItemText primary="Tags" />
                </ListItem>
                {/*<ListItem button component={Link} to="/settings">*/}
                {/*    <ListItemText primary="Settings" />*/}
                {/*</ListItem>*/}
                {/* Add more links as needed */}
            </List>
        </div>
    );
};

export default SideBar;
