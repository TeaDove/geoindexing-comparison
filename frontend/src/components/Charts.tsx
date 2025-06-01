import React, { useState, useEffect, useMemo, useRef } from 'react';
import { Line } from 'react-chartjs-2';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
    ChartData,
    ChartDataset
} from 'chart.js';
import * as errorBarPlugin from 'chartjs-chart-error-bars';
import type { Run } from '../types/index';
import { API_URL } from '../config';

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
);
ChartJS.register(errorBarPlugin);

interface Point {
    x: number;
    y: number;
}

interface Dataset {
    points: Point[];
    regressionPoints?: Point[];
}

interface StatsResponse {
    [task: string]: {
        [index: string]: Dataset;
    };
}

interface ChartsProps {
    selectedRunId: number | null;
    run?: Run;
}

const Charts: React.FC<ChartsProps> = ({ selectedRunId, run }) => {
    const [stats, setStats] = useState<StatsResponse | null>(null);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [retryCount, setRetryCount] = useState(0);
    const [fullscreenTask, setFullscreenTask] = useState<string | null>(null);
    const [showRegression, setShowRegression] = useState<Record<string, boolean>>({});
    const chartRefs = useRef<Record<string, React.MutableRefObject<any>>>({});

    // Generate consistent colors for indexes
    const colorMap = useMemo(() => new Map<string, string>(), []);

    const getColorForIndex = (index: string) => {
        if (!colorMap.has(index)) {
            const colorIndex = colorMap.size % 10;
            colorMap.set(index, [
                '#4C72B0',
                '#DD8452',
                '#55A868',
                '#C44E52',
                '#8172B3',
                '#937860',
                '#DA8BC3',
                '#8C8C8C',
                '#CCB974',
                '#64B5CD'
            ][colorIndex]);
        }
        return colorMap.get(index)!;
    };

    useEffect(() => {
        let isMounted = true;
        let intervalId: number;

        const fetchStats = async () => {
            if (!selectedRunId) return;

            try {
                const response = await fetch(`${API_URL}/runs/stats`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ runId: selectedRunId }),
                });

                if (!response.ok) {
                    throw new Error('Failed to fetch stats');
                }

                const data: StatsResponse = await response.json();
                if (isMounted) {
                    setStats(data);
                    setError(null);
                    setRetryCount(0);
                }
            } catch (err) {
                if (isMounted) {
                    console.error('Error fetching stats:', err);
                    setError(err instanceof Error ? err.message : 'Failed to fetch stats');
                }
            } finally {
                if (isMounted) {
                    setIsLoading(false);
                }
            }
        };

        const startPolling = () => {
            setIsLoading(true);
            fetchStats();

            // Only poll if run is not completed
            if (run && run.Status !== 'COMPLETED') {
                intervalId = window.setInterval(fetchStats, 1000);
            }
        };

        startPolling();

        return () => {
            isMounted = false;
            if (intervalId) {
                clearInterval(intervalId);
            }
        };
    }, [selectedRunId, run?.Status, retryCount]);

    const toggleRegression = (task: string) => {
        setShowRegression(prev => ({
            ...prev,
            [task]: !prev[task]
        }));
    };

    if (!selectedRunId) {
        return null;
    }

    if (isLoading && !stats) {
        return (
            <div className="charts-loading">
                Loading charts... {retryCount > 0 ? `(Retry ${retryCount}/3)` : ''}
            </div>
        );
    }

    if (error) {
        return (
            <div className="charts-error">
                <div>Error: {error}</div>
                {retryCount < 3 && (
                    <button
                        onClick={() => setRetryCount(prev => prev + 1)}
                        className="retry-button"
                    >
                        Retry
                    </button>
                )}
            </div>
        );
    }

    if (!stats || Object.keys(stats).length === 0) {
        return <div className="charts-empty">No data available</div>;
    }

    try {
        // Get all tasks from the stats
        const tasks = Object.keys(stats);

        return (
            <div className="charts-container">
                {tasks.map(task => {
                    const taskData = stats[task];
                    const indexes = Object.keys(taskData);
                    const hasRegressionPoints = indexes.some(index => {
                        const dataset = taskData[index];
                        const regressionPoints = dataset?.regressionPoints;
                        return regressionPoints !== undefined && regressionPoints.length > 0;
                    });

                    const datasets: ChartDataset<"line", Point[]>[] = indexes.flatMap(index => {
                        const dataset = taskData[index];
                        if (!dataset) return [];

                        const baseDataset: ChartDataset<"line", Point[]> = {
                            label: index,
                            data: dataset.points,
                            borderColor: getColorForIndex(index),
                            tension: 0.4,
                            cubicInterpolationMode: "monotone" as const,
                            showLine: true,
                            pointStyle: 'circle',
                        };

                        // If regression points are enabled and available, add them as a separate dataset
                        if (showRegression[task] && dataset.regressionPoints && dataset.regressionPoints.length > 0) {
                            const regressionDataset: ChartDataset<"line", Point[]> = {
                                label: `${index} (регрессия)`,
                                data: dataset.regressionPoints,
                                borderColor: getColorForIndex(index),
                                backgroundColor: getColorForIndex(index),
                                borderWidth: 2,
                                borderDash: [5, 5],
                                tension: 0.4,
                                cubicInterpolationMode: "monotone" as const,
                                showLine: true,
                                pointStyle: false,
                                fill: false,
                            };
                            return [baseDataset, regressionDataset];
                        }

                        return [baseDataset];
                    });

                    const chartData: ChartData<"line", Point[]> = {
                        datasets: datasets.map(dataset => {
                            const isRegression = typeof dataset.label === 'string' && dataset.label.includes('регрессия');
                            return {
                                ...dataset,
                                borderColor: dataset.borderColor + (isRegression ? '66' : ''), // Add 40% opacity for regression lines
                            };
                        })
                    };

                    const options = {
                        responsive: true,
                        plugins: {
                            legend: {
                                position: 'top' as const,
                                align: 'start' as const,
                            },
                            title: {
                                display: true,
                                text: `Задача: ${task}`
                            }
                        },
                        scales: {
                            x: {
                                type: 'linear' as const,
                                title: {
                                    display: true,
                                    text: 'Кол-во точек'
                                }
                            },
                            y: {
                                title: {
                                    display: true,
                                    text: 'Время выполнения (μs)'
                                }
                            }
                        }
                    };

                    // Fullscreen modal for this chart
                    const isFullscreen = fullscreenTask === task;

                    if (!chartRefs.current[task]) {
                        chartRefs.current[task] = React.createRef<any>();
                    }

                    return (
                        <React.Fragment key={task}>
                            <div className="chart-wrapper">
                                <div style={{ display: 'flex', justifyContent: 'flex-end', gap: 8, marginBottom: 8 }}>
                                    {hasRegressionPoints && (
                                        <button onClick={() => toggleRegression(task)} >
                                            {showRegression[task] ? '↗' : '⦨'}
                                        </button>
                                    )}
                                    <button onClick={() => setFullscreenTask(task)}>
                                        ⛶
                                    </button>
                                </div>
                                <Line ref={chartRefs.current[task]} data={chartData} options={options} />
                            </div>
                            {isFullscreen && (
                                <div className="fullscreen-modal" onClick={() => setFullscreenTask(null)}>
                                    <div className="fullscreen-chart" onClick={e => e.stopPropagation()}>
                                        <div style={{ display: 'flex', justifyContent: 'flex-end', gap: 8, marginBottom: 8 }}>
                                            {hasRegressionPoints && (
                                                <button onClick={() => toggleRegression(task)} >
                                                    {showRegression[task] ? '↗' : '⦨'}
                                                </button>
                                            )}
                                            <button onClick={() => setFullscreenTask(null)}>
                                                ⛶
                                            </button>
                                        </div>
                                        <Line data={chartData} options={options} height={600} />
                                    </div>
                                </div>
                            )}
                        </React.Fragment>
                    );
                })}
            </div>
        );
    } catch (err) {
        console.error('Error rendering charts:', err);
        return (
            <div className="charts-error">
                Error rendering charts. Please try again.
            </div>
        );
    }
};

export default Charts; 