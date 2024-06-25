// src/services/songApi.ts
import api from './api';

export const getSongs = () => api.get('/admin/songs');

