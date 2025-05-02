import React from 'react';
import { Link } from 'react-router-dom';
import '../App.css'; // Import common styles

const Visualizer: React.FC = () => {
    return (
        <div className="page-container visualizer-page">
            <nav>
                <Link to="/">Go to Chart</Link>
            </nav>
            <h1>Kepler.gl Visualization</h1>
            <div className="kepler-gl-container">
                <iframe
                    title="Kepler.gl Embed"
                    src="https://kepler.gl/demo" // Replace with your specific Kepler.gl URL if needed
                    width="100%"
                    height="800px" // Adjust height as needed
                    style={{ border: 'none' }}
                ></iframe>
            </div>
        </div>
    );
};

export default Visualizer; 