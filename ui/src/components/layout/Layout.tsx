import React, { ReactNode } from 'react';
import { Box } from '@mui/material';
import SideBar from './SideBar';
import Navbar from "./Navbar";

interface LayoutProps {
    children: ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
    return (
        <Box display="flex">

            <SideBar />
            <Box component="main" flexGrow={1} p={3}>
                <Navbar />
                {children}
            </Box>
        </Box>
    );
};

export default Layout;