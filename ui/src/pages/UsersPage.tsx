// src/pages/UsersPage.tsx
import React from 'react';
import Navbar from '../components/layout/Navbar';
import Sidebar from '../components/layout/SideBar';
import useUsers from '../hooks/useUsers';
import { CircularProgress, List, ListItem, ListItemText } from '@mui/material';

const UsersPage: React.FC = () => {
    const { users, loading } = useUsers();

    return (
        <div>
            <Navbar />
        <Sidebar />
        <main style={{ marginLeft: 240, padding: 16 }}>
    <h1>Users</h1>
    {loading ? (
        <CircularProgress />
    ) : (
        <List>
            {users.map((user: any) => (
                    <ListItem key={user.id}>
                    <ListItemText primary={user.username} secondary={user.email} />
    </ListItem>
    ))}
        </List>
    )}
    </main>
    </div>
);
};

export default UsersPage;
