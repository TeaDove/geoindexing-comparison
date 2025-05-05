import React, { useState, useEffect, useRef } from 'react';
import { Link } from 'react-router-dom';
import { API_URL, MAPBOX_TOKEN } from '../config';
import Notification, { NotificationMessage } from '../components/Notification';
import '../App.css';
import mapboxgl, { Map, LngLatLike, Marker } from 'mapbox-gl';
import 'mapbox-gl/dist/mapbox-gl.css';

const headers = {
    'Content-Type': 'application/json'
};

const Visualizer: React.FC = () => {
    const [indexes, setIndexes] = useState<{ info: { shortName: string, longName: string } }[]>([]);
    const [selectedIndex, setSelectedIndex] = useState<string>('');
    const [amount, setAmount] = useState<number>(1000);
    const [isLoadingGenerate, setIsLoadingGenerate] = useState(false);
    const [isLoadingPoints, setIsLoadingPoints] = useState(false);
    const [notification, setNotification] = useState<NotificationMessage | null>(null);

    // --- KNN State ---
    const [selectedPoint, setSelectedPoint] = useState<mapboxgl.LngLat | null>(null);
    const [knnN, setKnnN] = useState<number>(10);
    const [isLoadingKnn, setIsLoadingKnn] = useState(false);
    const [knnNeighborsGeoJson, setKnnNeighborsGeoJson] = useState<GeoJSON.FeatureCollection<GeoJSON.Point> | null>(null);

    // --- Mapbox State and Refs ---
    const mapContainerRef = useRef<HTMLDivElement>(null);
    const mapRef = useRef<Map | null>(null);
    const selectedMarkerRef = useRef<Marker | null>(null); // Ref for the selected point marker

    useEffect(() => {
        const fetchIndexes = async () => {
            const startTime = performance.now(); // Start timer
            let status = 500; // Default status
            let errorMsg: string | undefined;
            try {
                const response = await fetch(`${API_URL}/indexes`, { headers });
                status = response.status; // Get status
                if (!response.ok) {
                    errorMsg = await response.text() || 'Failed to fetch indexes';
                    throw new Error(errorMsg);
                }
                const data: { info: { shortName: string, longName: string } }[] = await response.json();
                setIndexes(data);
                if (data.length > 0 && data[0]?.info?.shortName) {
                    setSelectedIndex(data[0].info.shortName);
                }
            } catch (error) {
                console.error('Error fetching indexes:', error);
                if (error instanceof Error) {
                    if (!error.message.includes('Failed to fetch indexes')) {
                        errorMsg = error.message;
                    }
                }
                if (status !== 200 && !errorMsg) {
                    errorMsg = 'Caught an unknown error';
                }
            } finally {
                const endTime = performance.now(); // End timer
                const durationMs = endTime - startTime;
                showNotification(status, '/indexes', 'GET', errorMsg, durationMs); // Show notification with duration
            }
        };
        fetchIndexes();
    }, []);

    useEffect(() => {
        if (mapRef.current || !mapContainerRef.current) return;

        mapboxgl.accessToken = MAPBOX_TOKEN;
        const map = new Map({
            container: mapContainerRef.current,
            style: 'mapbox://styles/mapbox/light-v11',
            center: [37.6173, 55.7558],
            zoom: 11,
        });

        map.on('load', () => {
            mapRef.current = map;
            console.log('Mapbox map loaded.');

            // Source and Layer for initially loaded points
            map.addSource('points', {
                type: 'geojson',
                data: { type: 'FeatureCollection', features: [] }
            });
            map.addLayer({
                id: 'points-layer',
                type: 'circle',
                source: 'points',
                paint: { 'circle-radius': 4, 'circle-color': '#007cbf' }
            });

            // Source and Layer for KNN Neighbors
            map.addSource('knn-neighbors', {
                type: 'geojson',
                data: { type: 'FeatureCollection', features: [] } // Initially empty
            });
            map.addLayer({
                id: 'knn-neighbors-layer',
                type: 'circle',
                source: 'knn-neighbors',
                paint: { 'circle-radius': 5, 'circle-color': '#2ca02c' } // Green color
            });
        });

        // --- Map Click Handler ---
        map.on('click', (e) => {
            const coordinates = e.lngLat;
            console.log('Map clicked at:', coordinates);
            setSelectedPoint(coordinates);
            setKnnNeighborsGeoJson(null); // Clear previous KNN results on new click

            // Remove previous marker if it exists
            if (selectedMarkerRef.current) {
                selectedMarkerRef.current.remove();
            }

            // Add a new marker for the selected point (Red)
            const marker = new Marker({ color: '#d62728' })
                .setLngLat(coordinates)
                .addTo(map);
            selectedMarkerRef.current = marker; // Store the new marker instance
        });

        // Clean up on unmount
        return () => {
            map.remove();
            mapRef.current = null;
            selectedMarkerRef.current = null; // Clear marker ref on unmount
        };
    }, []);

    const showNotification = (status: number, endpoint: string, method: string, error?: string, durationMs?: number) => {
        console.log(`Notification: ${method} ${endpoint} -> ${status}${error ? ` Error: ${error}` : ''}${durationMs ? ` (${durationMs.toFixed(2)} ms)` : ''}`);
        setNotification({
            status,
            endpoint,
            method,
            error,
            durationMs,
            timestamp: Date.now(),
        });
    };

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        if (!selectedIndex) {
            showNotification(400, '/visualizer', 'POST', 'Please select an index.') // No duration for client-side validation
            return;
        }
        setIsLoadingGenerate(true);
        const startTime = performance.now();
        let status = 500;
        let errorMsg: string | undefined;
        try {
            const response = await fetch(`${API_URL}/visualizer`, {
                method: 'POST',
                headers,
                body: JSON.stringify({ index: selectedIndex, amount }),
            });
            status = response.status;
            if (!response.ok) {
                errorMsg = await response.text() || 'Failed to submit visualizer request';
                throw new Error(errorMsg);
            }
            console.log('Visualizer POST successful.');
        } catch (error) {
            console.error('Error submitting visualizer request:', error);
            if (error instanceof Error) {
                if (!error.message.includes('Failed to submit visualizer request')) {
                    errorMsg = error.message;
                }
            }
            if (status !== 200 && !errorMsg) {
                errorMsg = 'Caught an unknown error';
            }
        } finally {
            const endTime = performance.now();
            const durationMs = endTime - startTime;
            showNotification(status, '/visualizer', 'POST', errorMsg, durationMs);
            setIsLoadingGenerate(false);
        }
    };

    const handleLoadData = async () => {
        setIsLoadingPoints(true);
        setSelectedPoint(null);
        setKnnNeighborsGeoJson(null);
        if (selectedMarkerRef.current) {
            selectedMarkerRef.current.remove();
            selectedMarkerRef.current = null;
        }
        const startTime = performance.now();
        let status = 500;
        let errorMsg: string | undefined;
        let geoJsonData: GeoJSON.FeatureCollection<GeoJSON.Point> | null = null;
        try {
            const response = await fetch(`${API_URL}/visualizer/points`);
            status = response.status;
            if (!response.ok) {
                errorMsg = await response.text() || 'Failed to fetch points';
                throw new Error(errorMsg);
            }
            geoJsonData = await response.json();
            console.log("Fetched GeoJSON data:", geoJsonData);

        } catch (error) {
            console.error('Error loading points:', error);
            if (error instanceof Error) {
                if (!error.message.includes('Failed to fetch points')) {
                    errorMsg = error.message;
                }
            }
            if (status !== 200 && !errorMsg) {
                errorMsg = 'Caught an unknown error';
            }
        } finally {
            const endTime = performance.now();
            const durationMs = endTime - startTime;
            showNotification(status, '/visualizer/points', 'GET', errorMsg, durationMs);

            // Only update map if data fetch was successful
            if (status === 200 && geoJsonData && mapRef.current) {
                const source = mapRef.current.getSource('points') as mapboxgl.GeoJSONSource;
                if (source) {
                    source.setData(geoJsonData);
                    console.log('Updated Mapbox source with data.');
                    // No need for separate map success notification here, handled above
                } else {
                    console.error('Mapbox source "points" not found.');
                    showNotification(500, 'Mapbox', 'Source Update', 'Source not found');
                }
            } else if (status === 200 && geoJsonData && !mapRef.current) {
                console.error('Map instance not available when trying to update source.');
                showNotification(500, 'Mapbox', 'Source Update', 'Map not ready');
            }
            setIsLoadingPoints(false);
        }
    };

    const handleFindKnn = async () => {
        if (!selectedPoint) {
            showNotification(400, '/visualizer/knn', 'POST', 'No point selected on the map.');
            return;
        }
        setIsLoadingKnn(true);
        setKnnNeighborsGeoJson(null);
        const startTime = performance.now();
        let status = 500;
        let errorMsg: string | undefined;
        try {
            const payload = { point: { lon: selectedPoint.lng, lat: selectedPoint.lat }, n: knnN };
            const response = await fetch(`${API_URL}/visualizer/knn`, {
                method: 'POST',
                headers,
                body: JSON.stringify(payload),
            });
            status = response.status;
            if (!response.ok) {
                errorMsg = await response.text() || 'Failed to find KNN';
                throw new Error(errorMsg);
            }
            const neighborsData: GeoJSON.FeatureCollection<GeoJSON.Point> = await response.json();
            setKnnNeighborsGeoJson(neighborsData);
            console.log('KNN search successful, received neighbors:', neighborsData);

        } catch (error) {
            console.error('Error finding KNN:', error);
            if (error instanceof Error) {
                if (!error.message.includes('Failed to find KNN')) {
                    errorMsg = error.message;
                }
            }
            if (status !== 200 && !errorMsg) {
                errorMsg = 'Caught an unknown error';
            }
        } finally {
            const endTime = performance.now();
            const durationMs = endTime - startTime;
            showNotification(status, '/visualizer/knn', 'POST', errorMsg, durationMs);
            setIsLoadingKnn(false);
        }
    };

    // --- Effect to Update KNN Neighbors Layer ---
    useEffect(() => {
        if (mapRef.current) {
            const source = mapRef.current.getSource('knn-neighbors') as mapboxgl.GeoJSONSource;
            if (source) {
                // Use the state value, provide empty data if null
                source.setData(knnNeighborsGeoJson || { type: 'FeatureCollection', features: [] });
                console.log('Updated knn-neighbors source');
            }
        }
    }, [knnNeighborsGeoJson]); // Re-run when knnNeighborsGeoJson changes

    return (
        <div className="page-container visualizer-page">
            <nav>
                <Link to="/">Go to Chart</Link>
            </nav>
            <h1>Mapbox GL JS Visualization</h1>

            <form onSubmit={handleSubmit} className="visualizer-form">
                <div className="form-group">
                    <label htmlFor="index-select">Select Index:</label>
                    <select
                        id="index-select"
                        value={selectedIndex}
                        onChange={(e) => setSelectedIndex(e.target.value)}
                        disabled={indexes.length === 0 || isLoadingGenerate}
                    >
                        {indexes.map(index => (
                            index?.info?.shortName && (
                                <option key={index.info.shortName} value={index.info.shortName}>
                                    {index.info.longName} ({index.info.shortName})
                                </option>
                            )
                        ))}
                    </select>
                </div>
                <div className="form-group">
                    <label htmlFor="amount-input">Amount:</label>
                    <input
                        type="number"
                        id="amount-input"
                        value={amount}
                        onChange={(e) => setAmount(Number(e.target.value))}
                        min="1"
                        disabled={isLoadingGenerate}
                    />
                </div>
                <button type="submit" disabled={isLoadingGenerate} className="submit-button">
                    {isLoadingGenerate ? 'Generating...' : 'Generate Visualization Data'}
                </button>
            </form>

            <button
                onClick={handleLoadData}
                disabled={isLoadingPoints || isLoadingGenerate}
                className="load-data-button"
            >
                {isLoadingPoints ? 'Loading...' : 'Load Points from Backend'}
            </button>

            {/* --- KNN UI --- */}
            <div className="knn-controls form-group" style={{ marginTop: '15px', display: 'flex', alignItems: 'center', gap: '10px' }}>
                <label htmlFor="knn-n-input">Neighbors (N):</label>
                <input
                    type="number"
                    id="knn-n-input"
                    value={knnN}
                    onChange={(e) => setKnnN(Math.max(1, Number(e.target.value)))} // Ensure N is at least 1
                    min="1"
                    disabled={isLoadingKnn}
                    style={{ width: '60px' }}
                />
                <button
                    onClick={handleFindKnn} // We will define this function next
                    disabled={!selectedPoint || isLoadingKnn || isLoadingGenerate}
                    className="knn-button"
                >
                    {isLoadingKnn ? 'Finding...' : 'Find KNN'}
                </button>
                {selectedPoint && (
                    <span style={{ fontSize: '0.9em' }}>
                        Selected: {selectedPoint.lat.toFixed(4)}, {selectedPoint.lng.toFixed(4)}
                    </span>
                )}
            </div>
            {/* --- End KNN UI --- */}

            <Notification message={notification} />

            <div ref={mapContainerRef} className="mapbox-gl-container" style={{ height: '800px', width: '100%', marginTop: '20px' }} />

        </div>
    );
};

export default Visualizer; 