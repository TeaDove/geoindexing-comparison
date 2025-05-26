import React from 'react';
import { Outlet, useLocation } from 'react-router-dom';
import Header from './Header';

const TITLES: Record<string, string> = {
    '/charts': 'Чарты',
    '/visualizer': 'Визуализация',
    '/': 'Чарты',
};

const NAV_LINKS = [
    { to: '/charts', label: 'Чарты' },
    { to: '/visualizer', label: 'Визуализация' },
];

const Layout: React.FC = () => {
    const location = useLocation();
    const title = TITLES[location.pathname] || 'Чарты';

    return (
        <>
            <Header title={title} navLinks={NAV_LINKS} />
            <Outlet />
        </>
    );
};

export default Layout; 