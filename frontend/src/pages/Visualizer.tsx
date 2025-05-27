import React, { useState, useEffect, useRef } from 'react';
import { API_URL, MAPBOX_TOKEN } from '../config';
import '../App.css';
import mapboxgl, { Map, Marker } from 'mapbox-gl';
import 'mapbox-gl/dist/mapbox-gl.css';

const headers = {
    'Content-Type': 'application/json'
};

const Visualizer: React.FC = () => {
    const [indexes, setIndexes] = useState<{ info: { shortName: string, longName: string } }[]>([]);
    const [selectedIndex, setSelectedIndex] = useState<string>('');
    const [amount, setAmount] = useState<number>(10000);
    const [isLoadingGenerate, setIsLoadingGenerate] = useState(false);

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
    const selectedMarkerRef = useRef<Marker | null>(null);

    useEffect(() => {
        const fetchIndexes = async () => {
            try {
                const response = await fetch(`${API_URL}/indexes`, { headers });
                if (!response.ok) {
                    const error = await response.text() || 'Failed to fetch indexes';
                    throw new Error(error);
                }
                const data: { info: { shortName: string, longName: string } }[] = await response.json();
                setIndexes(data);
                if (data.length > 0 && data[0]?.info?.shortName) {
                    setSelectedIndex(data[0].info.shortName);
                }
            } catch (error) {
                console.error('Error fetching indexes:', error);
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

            // Set default selected point
            const defaultPoint = new mapboxgl.LngLat(37.609461, 55.728292);
            setSelectedPoint(defaultPoint);
            const marker = new Marker({ color: '#1E87F7' })
                .setLngLat(defaultPoint)
                .addTo(map);
            selectedMarkerRef.current = marker;

            // Source and Layer for initially loaded points
            map.addSource('points', {
                type: 'geojson',
                data: { type: 'FeatureCollection', features: [] }
            });
            map.addLayer({
                id: 'points-layer',
                type: 'circle',
                source: 'points',
                paint: { 'circle-radius': 4, 'circle-color': '#10B981' }
            });

            // Source and Layer for KNN Neighbors
            map.addSource('knn-neighbors', {
                type: 'geojson',
                data: { type: 'FeatureCollection', features: [] }
            });
            map.addLayer({
                id: 'knn-neighbors-layer',
                type: 'circle',
                source: 'knn-neighbors',
                paint: { 'circle-radius': 5, 'circle-color': '#FBBF24' }
            });

            // Source and Layer for Radius Search Results
            map.addSource('radius-search-results', {
                type: 'geojson',
                data: { type: 'FeatureCollection', features: [] }
            });
            map.addLayer({
                id: 'radius-search-layer',
                type: 'circle',
                source: 'radius-search-results',
                paint: { 'circle-radius': 5, 'circle-color': '#F7701E' }
            });

            map.addSource('radius-circle', {
                type: 'geojson',
                data: { type: 'FeatureCollection', features: [] },
            });
        });

        // --- Map Click Handler ---
        map.on('click', (e) => {
            const coordinates = e.lngLat;
            console.log('Map clicked at:', coordinates);
            setSelectedPoint(coordinates);
            setKnnNeighborsGeoJson(null);
            setRadiusSearchResultsGeoJson(null);

            if (selectedMarkerRef.current) {
                selectedMarkerRef.current.remove();
            }

            const marker = new Marker({ color: '#DC2626' })
                .setLngLat(coordinates)
                .addTo(map);
            selectedMarkerRef.current = marker;
        });

        // Clean up on unmount
        return () => {
            map.remove();
            mapRef.current = null;
            selectedMarkerRef.current = null; // Clear marker ref on unmount
        };
    }, []);

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        if (!selectedIndex) {
            console.error('Please select an index.');
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

        try {
            const response = await fetch(`${API_URL}/visualizer`, {
                method: 'POST',
                headers,
                body: JSON.stringify({ amount, index: selectedIndex }),
            });
            if (!response.ok) {
                const error = await response.text() || 'Failed to generate points';
                throw new Error(error);
            }
            const geoJsonData: GeoJSON.FeatureCollection<GeoJSON.Point> = await response.json();

            // Update map source if request was successful
            if (mapRef.current) {
                const pointsSource = mapRef.current.getSource('points') as mapboxgl.GeoJSONSource;
                if (pointsSource) {
                    pointsSource.setData(geoJsonData);
                }
                const knnSource = mapRef.current.getSource('knn-neighbors') as mapboxgl.GeoJSONSource;
                if (knnSource) knnSource.setData({ type: 'FeatureCollection', features: [] });
                const radiusSource = mapRef.current.getSource('radius-search-results') as mapboxgl.GeoJSONSource;
                if (radiusSource) radiusSource.setData({ type: 'FeatureCollection', features: [] });
            }
        } catch (error) {
            console.error('Error generating points:', error);
        } finally {
            setIsLoadingGenerate(false);
        }
    };

    const handleFindKnn = async () => {
        if (!selectedPoint) {
            console.error('No point selected on the map.');
            return;
        }
        setIsLoadingKnn(true);
        setKnnNeighborsGeoJson(null);
        try {
            const payload = { point: { lon: selectedPoint.lng, lat: selectedPoint.lat }, n: knnN };
            const response = await fetch(`${API_URL}/visualizer/knn`, {
                method: 'POST',
                headers,
                body: JSON.stringify(payload),
            });
            if (!response.ok) {
                const error = await response.text() || 'Failed to find KNN';
                throw new Error(error);
            }
            const neighborsData: GeoJSON.FeatureCollection<GeoJSON.Point> = await response.json();
            setKnnNeighborsGeoJson(neighborsData);
            console.log('KNN search successful, received neighbors:', neighborsData);
        } catch (error) {
            console.error('Error finding KNN:', error);
        } finally {
            setIsLoadingKnn(false);
        }
    };

    const handleRadiusSearch = async () => {
        if (!selectedPoint) {
            console.error('No point selected on the map.');
            return;
        }
        setIsLoadingRadius(true);
        setRadiusSearchResultsGeoJson(null);
        try {
            const earthRadiusMeters = 6378137;
            const latDiffDegrees = (radius / earthRadiusMeters) * (180 / Math.PI);
            const lonDiffDegrees = (radius / (earthRadiusMeters * Math.cos(Math.PI * selectedPoint.lat / 180))) * (180 / Math.PI);

            const payload = {
                bottomLeft: { lon: selectedPoint.lng - lonDiffDegrees, lat: selectedPoint.lat - latDiffDegrees },
                upperRight: { lon: selectedPoint.lng + lonDiffDegrees, lat: selectedPoint.lat + latDiffDegrees }
            };

            const response = await fetch(`${API_URL}/visualizer/bbox`, {
                method: 'POST',
                headers,
                body: JSON.stringify(payload),
            });
            if (!response.ok) {
                const error = await response.text() || 'Failed to find points in bounding box';
                throw new Error(error);
            }
            const radiusData: GeoJSON.FeatureCollection<GeoJSON.Point> = await response.json();
            setRadiusSearchResultsGeoJson(radiusData);
            console.log('Bounding box search successful, received points:', radiusData);
        } catch (error) {
            console.error('Error finding points in bounding box:', error);
        } finally {
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

    // --- Effect to Load Points on Page Load (after map is ready) ---
    useEffect(() => {
        if (!mapRef.current) return;
        const loadInitialPoints = async () => {
            try {
                const response = await fetch(`${API_URL}/visualizer/points`, {
                    method: 'GET',
                    headers,
                });
                if (!response.ok) {
                    const error = await response.text() || 'Failed to load initial points';
                    throw new Error(error);
                }
                const geoJsonData: GeoJSON.FeatureCollection<GeoJSON.Point> = await response.json();
                console.log('Initial points loaded successfully');

                // Update map with points if we have data
                if (mapRef.current && geoJsonData) {
                    const pointsSource = mapRef.current.getSource('points') as mapboxgl.GeoJSONSource;
                    if (pointsSource) {
                        pointsSource.setData(geoJsonData);
                    }
                }

            } catch (error) {
                console.error('Error loading initial points:', error);
            }
        };
        loadInitialPoints();
    }, [mapRef.current]);

    // --- Effect to Recreate Index on Backend When Selection Changes ---
    useEffect(() => {
        const recreateIndex = async () => {
            if (!selectedIndex) return; // Don't run if index is empty

            console.log(`Index selected: ${selectedIndex}. Sending request to backend.`);

            try {
                const response = await fetch(`${API_URL}/visualizer`, {
                    method: 'POST',
                    headers,
                    body: JSON.stringify({ index: selectedIndex }),
                });
                if (!response.ok) {
                    const error = await response.text() || `Failed to switch index to ${selectedIndex}`;
                    throw new Error(error);
                }
                const geoJsonData: GeoJSON.FeatureCollection<GeoJSON.Point> = await response.json();
                console.log(`Backend acknowledged index switch to ${selectedIndex} and returned points data`);

                // Clear KNN and radius search results, but DO NOT clear selectedPoint
                setKnnNeighborsGeoJson(null);
                setRadiusSearchResultsGeoJson(null);
                if (mapRef.current) {
                    const pointsSource = mapRef.current.getSource('points') as mapboxgl.GeoJSONSource;
                    if (pointsSource) {
                        if (geoJsonData) {
                            pointsSource.setData(geoJsonData);
                        } else {
                            pointsSource.setData({ type: 'FeatureCollection', features: [] });
                        }
                    }

                    const knnSource = mapRef.current.getSource('knn-neighbors') as mapboxgl.GeoJSONSource;
                    if (knnSource) knnSource.setData({ type: 'FeatureCollection', features: [] });
                    const radiusSource = mapRef.current.getSource('radius-search-results') as mapboxgl.GeoJSONSource;
                    if (radiusSource) radiusSource.setData({ type: 'FeatureCollection', features: [] });
                }

            } catch (error) {
                console.error('Error switching index:', error);
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
            <div style={{ display: 'flex', flexDirection: 'row', gap: '32px', alignItems: 'flex-start' }}>
                {/* Index select panel */}
                <fieldset
                    id="indexes"
                    style={{
                        minWidth: '280px',
                        background: 'var(--neutral-light)',
                        border: '2px solid var(--border-color)',
                        borderRadius: '8px',
                        padding: '24px 24px 16px 24px',
                        margin: '10px',
                        height: 'fit-content',
                    }}
                >
                    <legend>Индексы</legend>
                    {indexes.map(index => (
                        index?.info?.shortName && (
                            <div key={index.info.shortName} style={{ marginBottom: '8px' }}>
                                <input
                                    type="radio"
                                    id={`index-${index.info.shortName}`}
                                    name="index-select"
                                    value={index.info.shortName}
                                    checked={selectedIndex === index.info.shortName}
                                    onChange={(e) => setSelectedIndex(e.target.value)}
                                    disabled={isLoadingGenerate}
                                    style={{ marginRight: '8px' }}
                                />
                                <label htmlFor={`index-${index.info.shortName}`}>{index.info.longName} ({index.info.shortName})</label>
                            </div>
                        )
                    ))}
                </fieldset>

                {/* Controls and map */}
                <div className="visualizer-form-container">
                    <form onSubmit={handleSubmit} className="visualizer-form">
                        <div className="points-input" style={{ display: 'flex', alignItems: 'center', gap: '10px' }}>
                            <div>
                                <label htmlFor="amount-input">Кол-во точек:</label>
                                <input
                                    type="number"
                                    id="amount-input"
                                    value={amount}
                                    onChange={(e) => setAmount(Number(e.target.value))}
                                    min="1"
                                    disabled={isLoadingGenerate}
                                />
                            </div>
                            <div>
                                <label htmlFor="knn-n-input">Соседи (N):</label>
                                <input
                                    type="number"
                                    id="knn-n-input"
                                    value={knnN}
                                    onChange={(e) => setKnnN(Math.max(1, Number(e.target.value)))}
                                    min="1"
                                    disabled={isLoadingKnn}
                                />
                            </div>
                            <div>
                                <label htmlFor="radius-input">Радиус (m):</label>
                                <input
                                    type="number"
                                    id="radius-input"
                                    value={radius}
                                    onChange={(e) => setRadius(Math.max(1, Number(e.target.value)))}
                                    min="1"
                                    disabled={isLoadingRadius}
                                />
                            </div>
                            <button type="submit" disabled={isLoadingGenerate || isLoadingKnn} className="submit-button">
                                {isLoadingGenerate ? 'Generating...' : 'Загрузить данные'}
                            </button>
                            <button
                                onClick={handleFindKnn}
                                disabled={isLoadingKnn || isLoadingGenerate}
                                className="knn-button"
                            >
                                {isLoadingKnn ? 'Finding...' : 'Найти ближайщих соседей'}
                            </button>
                            <button
                                onClick={handleRadiusSearch}
                                disabled={isLoadingRadius || isLoadingGenerate}
                                className="radius-button"
                            >
                                {isLoadingRadius ? 'Searching...' : 'Найти в квадрате'}
                            </button>
                            {selectedPoint && (
                                <span style={{ fontSize: '0.9em' }}>
                                    Выбранная точка: {selectedPoint.lat.toFixed(4)}, {selectedPoint.lng.toFixed(4)}
                                </span>
                            )}
                        </div>
                    </form>

                    <div ref={mapContainerRef} className="mapbox-gl-container" style={{ height: '800px' }} />
                </div>
            </div>
        </div>
    );
};

export default Visualizer; 