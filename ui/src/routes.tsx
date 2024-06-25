import React from 'react';
import {Routes, Route, Outlet} from 'react-router-dom';
import DashboardPage from './pages/DashboardPage';
import UsersPage from './pages/UsersPage';
import StationsPage from './pages/StationsPage';
import SongsPage from './pages/SongsPage';
import TagsPage from './pages/TagsPage';
import Login from './pages/Login';
import Layout from './components/layout/Layout';
import ProtectedRoute from './components/ProtectedRoute';
import { AuthProvider } from './contexts/AuthContext';
import useAuth from "./hooks/useAuth";

const AppRoutes: React.FC = () => {
    const { user, loading } = useAuth();

    if (loading) {
        return <div>Loading...</div>; // Or a loading spinner
    }

    return (
        <AuthProvider>
            <Routes>
                <Route path="/login" element={<Login />} />
                <Route path="/" element={
                    <ProtectedRoute user={user}>
                        <Layout>
                            <DashboardPage />
                        </Layout>
                    </ProtectedRoute>
                } />
                <Route path="/users" element={
                    <ProtectedRoute user={user}>
                        <Layout>
                            <UsersPage />
                        </Layout>
                    </ProtectedRoute>
                } />
                <Route path="/stations" element={
                    <ProtectedRoute user={user}>
                        <Layout>
                            <StationsPage />
                        </Layout>
                    </ProtectedRoute>
                } />
                <Route path="/stations/:stationId/songs" element={
                    <ProtectedRoute user={user}>
                        <Layout>
                            <SongsPage />
                        </Layout>
                    </ProtectedRoute>
                } />
                <Route path="/songs" element={
                    <ProtectedRoute user={user}>
                        <Layout>
                            <SongsPage />
                        </Layout>
                    </ProtectedRoute>
                } />
                <Route path="/tags" element={
                    <ProtectedRoute user={user}>
                        <Layout>
                            <TagsPage />
                        </Layout>
                    </ProtectedRoute>
                } />
            </Routes>
        </AuthProvider>
    );
};

export default AppRoutes;
