import { useState, useEffect } from 'react';
import { getTags, createTag, updateTag, deleteTag } from '../services/tagApi';

export interface Tag {
    id: number;
    name: string;
}

const useTags = () => {
    const [tags, setTags] = useState<Tag[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchTags = async () => {
            try {
                const response = await getTags();
                setTags(response.data);
            } catch (error) {
                setError('Failed to fetch tags');
            } finally {
                setLoading(false);
            }
        };

        fetchTags();
    }, []);

    const addTag = async (tag: { name: string }) => {
        try {
            const response = await createTag(tag);
            setTags((prevTags) => [...prevTags, response.data]);
        } catch (error) {
            setError('Failed to add tag');
        }
    };

    const updateExistingTag = async (id: number, tag: { name: string }) => {
        try {
            await updateTag(id, tag);
            setTags((prevTags) => prevTags.map((t) => (t.id === id ? { ...t, name: tag.name } : t)));
        } catch (error) {
            setError('Failed to update tag');
        }
    };

    const removeTag = async (id: number) => {
        try {
            await deleteTag(id);
            setTags((prevTags) => prevTags.filter((tag) => tag.id !== id));
        } catch (error) {
            setError('Failed to delete tag');
        }
    };

    return {
        tags,
        loading,
        error,
        addTag,
        updateExistingTag,
        removeTag,
    };
};

export default useTags;
