import { useState, useEffect } from 'react';
import { getStations, createStation, updateStation, deleteStation, getTags } from '../services/stationApi';
import { Tag } from './useTags';

interface Station {
    id: number;
    name: string;
    tags: string[];
}

const useStations = () => {
    const [stations, setStations] = useState<Station[]>([]);
    const [tags, setTags] = useState<Tag[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchStationsAndTags = async () => {
            try {
                const [stationsResponse, tagsResponse] = await Promise.all([getStations(), getTags()]);
                setStations(stationsResponse.data);
                setTags(tagsResponse.data);
            } catch (error) {
                setError('Failed to fetch stations or tags');
            } finally {
                setLoading(false);
            }
        };

        fetchStationsAndTags();
    }, []);

    const addStation = async (name: string, tags: string[]) => {
        try {
            const response = await createStation({ name, tags });
            setStations((prevStations) => [...prevStations, response.data]);
        } catch (error) {
            setError('Failed to add station');
        }
    };

    const updateExistingStation = async (id: number, name: string, tags: string[]) => {
        try {
            const response = await updateStation(id, { name, tags });
            setStations((prevStations) => prevStations.map((station) => (station.id === id ? response.data : station)));
        } catch (error) {
            setError('Failed to update station');
        }
    };

    const removeStation = async (id: number) => {
        try {
            await deleteStation(id);
            setStations((prevStations) => prevStations.filter((station) => station.id !== id));
        } catch (error) {
            setError('Failed to delete station');
        }
    };

    return {
        stations,
        tags,
        loading,
        error,
        addStation,
        updateExistingStation,
        removeStation,
    };
};

export default useStations;
