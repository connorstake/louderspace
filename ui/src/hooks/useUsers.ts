// src/hooks/useUsers.ts
import { useEffect, useState } from 'react';
import { getUsers } from '../services/userApi';

const useUsers = () => {
    const [users, setUsers] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        getUsers()
            .then(response => {
                setUsers(response.data);
                setLoading(false);
            })
            .catch(error => {
                console.error('Failed to fetch users', error);
                setLoading(false);
            });
    }, []);

    return { users, loading };
};

export default useUsers;
