import React from 'react';
import { Link } from 'react-router-dom';

const Visualizer: React.FC = () => {
    return (
        <div>
            <h1>Visualizer</h1>
            <p>This page will contain the visualizer.</p>
            <Link to="/">Go to Chart</Link>
        </div>
    );
};

export default Visualizer; 