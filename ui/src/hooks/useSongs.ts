import { useState, useEffect } from 'react';
import { getSongs, getSongsByStation, createSong, updateSong, deleteSong, getTags } from '../services/songApi';
import { Tag } from './useTags';
import { useParams } from 'react-router-dom';
import useAuth from "./useAuth";

interface Song {
    id: number;
    title: string;
    artist: string;
    genre: string;
    suno_id: string;
    tags: Tag[];
}

const useSongs = () => {
    const { user, loading: authLoading } = useAuth();
    const { stationId } = useParams<{ stationId?: string }>();
    const [songs, setSongs] = useState<Song[]>([]);
    const [tags, setTags] = useState<Tag[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    console.log(user)
    useEffect(() => {
        const userId = user ? user.id : undefined;
        const fetchSongsAndTags = async () => {
            if (authLoading || (stationId && !user)) {
                // Wait until authLoading is false and user is loaded when stationId is present
                return;
            }

            setLoading(true);
            try {
                let songsResponse;
                if (stationId) {
                    songsResponse = await getSongsByStation(stationId, userId);
                } else {
                    songsResponse = await getSongs();
                }
                const tagsResponse = await getTags();
                setSongs(songsResponse.data);
                setTags(tagsResponse.data);
            } catch (error) {
                setError('Failed to fetch songs or tags');
            } finally {
                setLoading(false);
            }
        };

        fetchSongsAndTags();
    }, [stationId, user, authLoading]);

    const addSong = async (song: { title: string, artist: string, genre: string, suno_id: string, tags: string[] }) => {
        try {
            const response = await createSong(song);
            setSongs((prevSongs) => [...prevSongs, response.data]);
        } catch (error) {
            setError('Failed to add song');
        }
    };

    const updateExistingSong = async (id: number, song: { title: string, artist: string, genre: string, suno_id: string, tags: string[] }) => {
        try {
            const response = await updateSong(id, song);
            setSongs((prevSongs) => prevSongs.map((s) => (s.id === id ? response.data : s)));
        } catch (error) {
            setError('Failed to update song');
        }
    };

    const removeSong = async (id: number) => {
        try {
            await deleteSong(id);
            setSongs((prevSongs) => prevSongs.filter((song) => song.id !== id));
        } catch (error) {
            setError('Failed to delete song');
        }
    };

    return {
        songs,
        tags,
        loading,
        error,
        addSong,
        updateExistingSong,
        removeSong,
    };
};

export default useSongs;
