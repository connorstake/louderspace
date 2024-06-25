// src/services/userApi.ts
import api from './api';

export const getUsers = () => api.get('/admin/users');

// You can add more user-related API calls here
