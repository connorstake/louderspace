import api from './api';

// API calls related to tags
export const getTags = () => api.get('/admin/tags');
export const createTag = (tag: { name: string }) => api.post('/admin/tags', tag);
export const updateTag = (id: number, tag: { name: string }) => api.put(`/admin/tags/${id}`, tag);
export const deleteTag = (id: number) => api.delete(`/admin/tags/${id}`);
