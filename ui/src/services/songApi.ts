import api from './api';

// API calls related to songs
export const getSongs = () => api.get('/songs');
export const getSongsByStation = (stationId: string, userId?: number) => {
    const url = userId ? `/stations/${stationId}/songs?user_id=${userId}` : `/stations/${stationId}/songs`;
    return api.get(url);
};
export const createSong = (song: { title: string, artist: string, genre: string, suno_id: string, tags: string[] }) => api.post('/admin/songs', song);
export const updateSong = (id: number, song: { title: string, artist: string, genre: string, suno_id: string, tags: string[] }) => api.put(`/admin/songs/${id}`, song);
export const deleteSong = (id: number) => api.delete(`/admin/songs/${id}`);
export const getTags = () => api.get('/admin/tags');
