import React from 'react';
import { Link } from 'react-router-dom';

interface NavLink {
    to: string;
    label: string;
}

interface HeaderProps {
    title: string;
    navLinks?: NavLink[];
    backLink?: string;
    backText?: string;
}

const Header: React.FC<HeaderProps> = ({ title, navLinks, backLink, backText }) => (
    <header style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', padding: '16px 16px 0 16px', marginBottom: '16px' }}>
        <nav style={{ display: 'flex', gap: '16px' }}>
            {navLinks && navLinks.map(link => (
                <Link key={link.to} to={link.to}>{link.label}</Link>
            ))}
            {!navLinks && backLink && <Link to={backLink}>{backText || 'Назад'}</Link>}
        </nav>
        <h1 style={{ margin: 0 }}>{title}</h1>
    </header>
);

export default Header; 