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
    const [amount, setAmount] = useState<number>(10000);
    const [isLoadingGenerate, setIsLoadingGenerate] = useState(false);
    const [isLoadingPoints, setIsLoadingPoints] = useState(false);
    const [notification, setNotification] = useState<NotificationMessage | null>(null);

    // --- KNN State ---
    const [selectedPoint, setSelectedPoint] = useState<mapboxgl.LngLat | null>(null);
    const [knnN, setKnnN] = useState<number>(100);
    const [isLoadingKnn, setIsLoadingKnn] = useState(false);
    const [knnNeighborsGeoJson, setKnnNeighborsGeoJson] = useState<GeoJSON.FeatureCollection<GeoJSON.Point> | null>(null);

    // --- Radius Search State ---
    const [radius, setRadius] = useState<number>(1000); // Default radius (e.g., 1000 meters)
    const [isLoadingRadius, setIsLoadingRadius] = useState(false);
    const [radiusSearchResultsGeoJson, setRadiusSearchResultsGeoJson] = useState<GeoJSON.FeatureCollection<GeoJSON.Point> | null>(null);

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

            // Source and Layer for Radius Search Results
            map.addSource('radius-search-results', {
                type: 'geojson',
                data: { type: 'FeatureCollection', features: [] } // Initially empty
            });
            map.addLayer({
                id: 'radius-search-layer',
                type: 'circle',
                source: 'radius-search-results',
                paint: { 'circle-radius': 5, 'circle-color': '#ff7f0e' } // Orange color
            });
        });

        // --- Map Click Handler ---
        map.on('click', (e) => {
            const coordinates = e.lngLat;
            console.log('Map clicked at:', coordinates);
            setSelectedPoint(coordinates);
            setKnnNeighborsGeoJson(null); // Clear previous KNN results
            setRadiusSearchResultsGeoJson(null); // Clear previous Radius results

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

    const fetchAndLoadPoints = async (mapInstance: Map | null) => {
        if (!mapInstance) {
            console.error("Map instance not available for fetchAndLoadPoints");
            showNotification(500, 'Map Operation', 'Source Update', 'Map not ready');
            return false;
        }

        const startTime = performance.now();
        let status = 500;
        let errorMsg: string | undefined;
        let success = false;

        try {
            const response = await fetch(`${API_URL}/visualizer/points`);
            status = response.status;
            if (!response.ok) {
                errorMsg = await response.text() || 'Failed to fetch points';
                throw new Error(errorMsg);
            }
            const geoJsonData: GeoJSON.FeatureCollection<GeoJSON.Point> = await response.json();
            console.log("Fetched GeoJSON data for points layer:", geoJsonData);

            // Update map source
            const source = mapInstance.getSource('points') as mapboxgl.GeoJSONSource;
            if (source) {
                source.setData(geoJsonData);
                console.log('Updated Mapbox source "points".');
                success = true;
            } else {
                console.error('Mapbox source "points" not found during update.');
                errorMsg = 'Source "points" not found'; // Set error message for notification
                status = 500; // Ensure status reflects the source update error
            }

        } catch (error) {
            console.error('Error loading points:', error);
            if (error instanceof Error && !error.message.includes('Failed to fetch points')) {
                errorMsg = error.message;
            }
            if (status < 400) status = 500; // Ensure error status if catch block is hit
            if (!errorMsg) errorMsg = 'Caught an unknown error fetching points';
        } finally {
            const endTime = performance.now();
            const durationMs = endTime - startTime;
            // Notify about the point fetch/load attempt
            showNotification(status, '/visualizer/points', 'GET', errorMsg, durationMs);
        }
        return success;
    };

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        if (!selectedIndex) {
            showNotification(400, '/visualizer', 'POST', 'Please select an index.')
            return;
        }
        setIsLoadingGenerate(true);
        // Clear previous points and KNN selection when generating new data
        setSelectedPoint(null);
        setKnnNeighborsGeoJson(null);
        setRadiusSearchResultsGeoJson(null); // Clear radius search results
        if (selectedMarkerRef.current) {
            selectedMarkerRef.current.remove();
            selectedMarkerRef.current = null;
        }
        // Also clear the main points layer
        if (mapRef.current) {
            const source = mapRef.current.getSource('points') as mapboxgl.GeoJSONSource;
            if (source) {
                source.setData({ type: 'FeatureCollection', features: [] });
            }
        }

        const startTime = performance.now();
        let status = 500;
        let errorMsg: string | undefined;
        let geoJsonData: GeoJSON.FeatureCollection<GeoJSON.Point> | null = null;
        try {
            const response = await fetch(`${API_URL}/visualizer`, {
                method: 'POST',
                headers,
                body: JSON.stringify({ index: selectedIndex, amount }),
            });
            status = response.status;
            if (!response.ok) {
                errorMsg = await response.text() || 'Failed to generate/load data';
                throw new Error(errorMsg);
            }
            geoJsonData = await response.json(); // Expect GeoJSON directly from POST
            console.log('Visualizer POST successful, received GeoJSON:', geoJsonData);

        } catch (error) {
            console.error('Error submitting visualizer request:', error);
            if (error instanceof Error) {
                if (!error.message.includes('Failed to generate/load data')) {
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

            // Update map source if request was successful
            if (status === 200 && geoJsonData && mapRef.current) {
                const source = mapRef.current.getSource('points') as mapboxgl.GeoJSONSource;
                if (source) {
                    source.setData(geoJsonData);
                    console.log('Updated Mapbox source with generated data.');
                } else {
                    console.error('Mapbox source "points" not found.');
                    showNotification(500, 'Mapbox', 'Source Update', 'Source not found after generate');
                }
            } else if (status === 200 && geoJsonData && !mapRef.current) {
                console.error('Map instance not available when trying to update source after generate.');
                showNotification(500, 'Mapbox', 'Source Update', 'Map not ready after generate');
            }

            setIsLoadingGenerate(false);
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

    // --- Handle Radius Search ---
    const handleRadiusSearch = async () => {
        if (!selectedPoint) {
            showNotification(400, '/visualizer/range-search', 'POST', 'No point selected on the map.');
            return;
        }
        setIsLoadingRadius(true);
        setRadiusSearchResultsGeoJson(null); // Clear previous results visually
        const startTime = performance.now();
        let status = 500;
        let errorMsg: string | undefined;
        try {
            const payload = {
                point: { lon: selectedPoint.lng, lat: selectedPoint.lat },
                radius: radius // Use radius state
            };
            const response = await fetch(`${API_URL}/visualizer/range-search`, {
                method: 'POST',
                headers,
                body: JSON.stringify(payload),
            });
            status = response.status;
            if (!response.ok) {
                errorMsg = await response.text() || 'Failed to find points in range';
                throw new Error(errorMsg);
            }
            const radiusData: GeoJSON.FeatureCollection<GeoJSON.Point> = await response.json();
            setRadiusSearchResultsGeoJson(radiusData); // Update state, triggering useEffect
            console.log('Range search successful, received points:', radiusData);

        } catch (error) {
            console.error('Error finding points in range:', error);
            if (error instanceof Error) {
                if (!error.message.includes('Failed to find points in range')) {
                    errorMsg = error.message;
                }
            }
            if (status !== 200 && !errorMsg) {
                errorMsg = 'Caught an unknown error during range search';
            }
        } finally {
            const endTime = performance.now();
            const durationMs = endTime - startTime;
            showNotification(status, '/visualizer/range-search', 'POST', errorMsg, durationMs);
            setIsLoadingRadius(false);
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

    // --- Effect to Recreate Index on Backend When Selection Changes ---
    useEffect(() => {
        const recreateIndex = async () => {
            if (!selectedIndex) return; // Don't run if index is empty

            console.log(`Index selected: ${selectedIndex}. Sending request to backend.`);

            const startTime = performance.now();
            let status = 500;
            let errorMsg: string | undefined;

            try {
                const response = await fetch(`${API_URL}/visualizer`, {
                    method: 'POST',
                    headers,
                    body: JSON.stringify({ index: selectedIndex }),
                });
                status = response.status;
                if (!response.ok) {
                    errorMsg = await response.text() || `Failed to switch index to ${selectedIndex}`;
                    throw new Error(errorMsg);
                }
                console.log(`Backend acknowledged index switch to ${selectedIndex}`);
                // Clear points and search results
                setSelectedPoint(null);
                setKnnNeighborsGeoJson(null);
                setRadiusSearchResultsGeoJson(null); // Now this exists
                if (selectedMarkerRef.current) {
                    selectedMarkerRef.current.remove();
                    selectedMarkerRef.current = null;
                }
                if (mapRef.current) {
                    const pointsSource = mapRef.current.getSource('points') as mapboxgl.GeoJSONSource;
                    if (pointsSource) pointsSource.setData({ type: 'FeatureCollection', features: [] });
                    const knnSource = mapRef.current.getSource('knn-neighbors') as mapboxgl.GeoJSONSource;
                    if (knnSource) knnSource.setData({ type: 'FeatureCollection', features: [] });
                    const radiusSource = mapRef.current.getSource('radius-search-results') as mapboxgl.GeoJSONSource;
                    // Need to add this source/layer first, will do later
                    // if (radiusSource) radiusSource.setData({ type: 'FeatureCollection', features: [] }); 
                }

            } catch (error) {
                console.error('Error switching index:', error);
                if (error instanceof Error) {
                    if (!error.message.includes('Failed to switch index')) { // Simpler check
                        errorMsg = error.message;
                    }
                }
                if (status !== 200 && !errorMsg) {
                    errorMsg = 'Caught an unknown error during index switch';
                }
            } finally {
                const endTime = performance.now();
                const durationMs = endTime - startTime;
                showNotification(status, `/visualizer (index: ${selectedIndex})`, 'POST', errorMsg, durationMs);
            }
        };

        recreateIndex();
    }, [selectedIndex]);

    // --- Effect to Update Radius Search Layer ---
    useEffect(() => {
        if (mapRef.current) {
            const source = mapRef.current.getSource('radius-search-results') as mapboxgl.GeoJSONSource;
            if (source) {
                source.setData(radiusSearchResultsGeoJson || { type: 'FeatureCollection', features: [] });
                console.log('Updated radius-search-results source');
            } else {
                // Source might not exist yet on initial load, this is okay
                // console.log('Radius search source not found yet.'); 
            }
        }
    }, [radiusSearchResultsGeoJson]);

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
                <button type="submit" disabled={isLoadingGenerate || isLoadingKnn} className="submit-button">
                    {isLoadingGenerate ? 'Generating...' : 'Generate and Load Data'}
                </button>
            </form>

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
                    onClick={handleFindKnn}
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

            {/* --- Radius Search UI --- */}
            <div className="radius-controls form-group" style={{ marginTop: '10px', display: 'flex', alignItems: 'center', gap: '10px' }}>
                <label htmlFor="radius-input">Radius (m):</label>
                <input
                    type="number"
                    id="radius-input"
                    value={radius}
                    onChange={(e) => setRadius(Math.max(1, Number(e.target.value)))} // Ensure radius is positive
                    min="1"
                    disabled={isLoadingRadius}
                    style={{ width: '80px' }}
                />
                <button
                    onClick={handleRadiusSearch}
                    disabled={!selectedPoint || isLoadingRadius || isLoadingGenerate}
                    className="radius-button"
                >
                    {isLoadingRadius ? 'Searching...' : 'Search Radius'}
                </button>
            </div>
            {/* --- End Radius Search UI --- */}

            <Notification message={notification} />

            <div ref={mapContainerRef} className="mapbox-gl-container" style={{ height: '800px', width: '100%', marginTop: '20px' }} />

        </div>
    );
};

export default Visualizer; 