// src/routes.tsx
import React from 'react';
import { Routes, Route } from 'react-router-dom';
import DashboardPage from './pages/DashboardPage';
import UsersPage from './pages/UsersPage';
import StationsPage from './pages/StationsPage';
import SongsPage from './pages/SongsPage';
import Layout from './components/layout/Layout';
// import SettingsPage from './pages/SettingsPage';
// import NotFoundPage from './pages/NotFoundPage';

const AppRoutes: React.FC = () => {
    return (
        <Layout>
            <Routes>
                <Route path="/" element={<DashboardPage />} />
                <Route path="/users" element={<UsersPage />} />
                <Route path="/stations" element={<StationsPage />} />
                <Route path="/stations/:stationId/songs" element={<SongsPage />} />
                <Route path="/songs" element={<SongsPage />} />
                {/*<Route path="/settings" element={<SettingsPage />} />*/}
                {/*<Route path="*" element={<NotFoundPage />} />*/}
            </Routes>
        </Layout>
    );
};

export default AppRoutes;
