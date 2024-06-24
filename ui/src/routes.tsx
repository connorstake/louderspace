// src/routes.tsx
import React from 'react';
import { Routes, Route } from 'react-router-dom';
import DashboardPage from './pages/DashboardPage';
import UsersPage from './pages/UsersPage';
import StationsPage from './pages/StationsPage';
import SongsPage from './pages/SongsPage';
// import SettingsPage from './pages/SettingsPage';
// import NotFoundPage from './pages/NotFoundPage';

const AppRoutes: React.FC = () => {
    return (
        <Routes>
            <Route path="/" element={<DashboardPage />} />
            <Route path="/users" element={<UsersPage />} />
            <Route path="/stations" element={<StationsPage />} />
            <Route path="/songs" element={<SongsPage />} />
            {/*<Route path="/settings" element={<SettingsPage />} />*/}
            {/*<Route path="*" element={<NotFoundPage />} />*/}
        </Routes>
    );
};

export default AppRoutes;
