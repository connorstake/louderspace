// src/services/authApi.ts
import api from './api';

export const login = (username: string, password: string) =>
    api.post('/login', { username, password });

export const register = (username: string, password: string, email: string, role: string) =>
    api.post('/register', { username, password, email, role });

