// src/services/api.ts
import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8080', // Adjust the base URL as needed
});

export const getUsers = () => api.get('/users');
export const getSongs = () => api.get('/songs');
// Add more API calls as needed

export default api;
