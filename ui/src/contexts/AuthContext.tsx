import React, { createContext, useState, useEffect, ReactNode, useContext } from 'react';
import useAuth from '../hooks/useAuth';

interface User {
    id: number;
    username: string;
    email: string;
    role: string;
}



interface AuthContextType {
    user: User | null;
    login: (username: string, password: string) => Promise<void>;
    register: (username: string, password: string, email: string, role: string) => Promise<void>;
    logout: () => void;
    loading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const auth = useAuth();

    console.log(auth);
    return <AuthContext.Provider value={auth}>{!auth.loading && children}</AuthContext.Provider>;
};

export const useAuthContext = (): AuthContextType => {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};
