// src/hooks/useAuth.ts
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { login as loginApi, register as registerApi } from '../services/authApi';
import api from '../services/api';

interface User {
    id: number;
    username: string;
    email: string;
    role: string;
}

const useAuth = () => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
            api.get('/me')
                .then(response => {
                    setUser(response.data);
                    setLoading(false); // Move this inside the success path
                })
                .catch(() => {
                    localStorage.removeItem('token');
                    setLoading(false); // Ensure this is set even on error
                });
        } else {
            setLoading(false); // Ensure this is set if there's no token
        }
    }, []);

    const login = async (username: string, password: string) => {
        const response = await loginApi(username, password);
        localStorage.setItem('token', response.data.token);
        api.defaults.headers.common['Authorization'] = `Bearer ${response.data.token}`;
        setUser(response.data.user);
        navigate('/');
    };

    const register = async (username: string, password: string, email: string, role: string) => {
        const response = await registerApi(username, password, email, role);
        localStorage.setItem('token', response.data.token);
        api.defaults.headers.common['Authorization'] = `Bearer ${response.data.token}`;
        setUser(response.data.user);
        navigate('/');
    };

    const logout = () => {
        localStorage.removeItem('token');
        delete api.defaults.headers.common['Authorization'];
        setUser(null);
        navigate('/login');
    };

    return { user, loading, login, register, logout };
};

export default useAuth;
