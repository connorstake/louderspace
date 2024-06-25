import api from './api';

// API calls related to stations
export const getStations = () => api.get('/stations');
export const createStation = (station: { name: string, tags: string[] }) => api.post('admin/stations', station);
export const updateStation = (id: number, station: { name: string, tags: string[] }) => api.put(`/admin/stations/${id}`, station);
export const deleteStation = (id: number) => api.delete(`/admin/stations/${id}`);
export const getTags = () => api.get('/admin/tags');
