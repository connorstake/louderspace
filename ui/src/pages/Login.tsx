// src/pages/LoginPage.tsx
import React, { useState, useEffect } from 'react';
import { TextField, Button, Container, Typography, CircularProgress } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import useAuth from "../hooks/useAuth";

const LoginPage: React.FC = () => {
    const { login, user, loading } = useAuth();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    useEffect(() => {
        if (!loading && user) {
            navigate('/');
        }
    }, [user, loading, navigate]);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            await login(username, password);
        } catch (err) {
            setError('Invalid username or password');
        }
    };

    if (loading) {
        return (
            <Container maxWidth="xs">
                <CircularProgress />
            </Container>
        );
    }

    return (
        <Container maxWidth="xs">
            <Typography variant="h4" component="h1" gutterBottom>
                Login
            </Typography>
            <form onSubmit={handleSubmit}>
                <TextField
                    label="Username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    fullWidth
                    margin="normal"
                />
                <TextField
                    label="Password"
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    fullWidth
                    margin="normal"
                />
                {error && <Typography color="error">{error}</Typography>}
                <Button type="submit" variant="contained" color="primary" fullWidth>
                    Login
                </Button>
            </form>
        </Container>
    );
};

export default LoginPage;
