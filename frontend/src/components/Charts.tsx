import React, { useState, useEffect, useMemo } from 'react';
import { Line } from 'react-chartjs-2';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
} from 'chart.js';
import * as errorBarPlugin from 'chartjs-chart-error-bars';
import type { Point, Run } from '../types/index';
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

interface ChartsProps {
    selectedRunId: number | null;
    run?: Run;
}

const Charts: React.FC<ChartsProps> = ({ selectedRunId, run }) => {
    const [points, setPoints] = useState<Point[]>([]);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [retryCount, setRetryCount] = useState(0);
    const [fullscreenTask, setFullscreenTask] = useState<string | null>(null);

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

        const fetchPoints = async () => {
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
                    throw new Error('Failed to fetch points');
                }

                const data = await response.json();
                if (isMounted) {
                    setPoints(data);
                    setError(null);
                    setRetryCount(0);
                }
            } catch (err) {
                if (isMounted) {
                    console.error('Error fetching points:', err);
                    setError(err instanceof Error ? err.message : 'Failed to fetch points');

                    // Retry up to 3 times with exponential backoff
                    if (retryCount < 3) {
                        setTimeout(() => {
                            setRetryCount(prev => prev + 1);
                            fetchPoints();
                        }, Math.pow(2, retryCount) * 1000);
                    }
                }
            } finally {
                if (isMounted) {
                    setIsLoading(false);
                }
            }
        };

        const startPolling = () => {
            setIsLoading(true);
            fetchPoints();

            // Only poll if run is not completed
            if (run && run.Status !== 'COMPLETED') {
                intervalId = window.setInterval(fetchPoints, 1000);
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

    if (!selectedRunId) {
        return null;
    }

    if (isLoading && points.length === 0) {
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

    if (points.length === 0) {
        return <div className="charts-empty">No data available</div>;
    }

    try {
        // Group points by task
        const tasks = Array.from(new Set(points.map(p => p.task)));

        return (
            <div className="charts-container">
                {tasks.map(task => {
                    // Filter points for this task
                    const taskPoints = points.filter(p => p.task === task);

                    // Group points by index
                    const datasets = taskPoints.reduce((acc, point) => {
                        if (!acc[point.index]) {
                            acc[point.index] = {
                                label: point.index,
                                data: [],
                                borderColor: getColorForIndex(point.index),
                                tension: 0.4,
                                cubicInterpolationMode: 'monotone',
                                showLine: true,
                                pointStyle: 'circle',
                                errorBarWhiskerColor: getColorForIndex(point.index),
                            };
                        }
                        acc[point.index].data.push({ x: point.x, y: point.y });
                        return acc;
                    }, {} as Record<string, any>);

                    const chartData = {
                        datasets: Object.values(datasets)
                    };

                    const options = {
                        responsive: true,
                        plugins: {
                            legend: {
                                position: 'top' as const,
                            },
                            title: {
                                display: true,
                                text: `Task: ${task}`
                            }
                        },
                        scales: {
                            x: {
                                type: 'linear' as const,
                                title: {
                                    display: true,
                                    text: 'Points'
                                }
                            },
                            y: {
                                title: {
                                    display: true,
                                    text: 'Time (μs)'
                                }
                            }
                        }
                    };

                    // Fullscreen modal for this chart
                    const isFullscreen = fullscreenTask === task;

                    return (
                        <React.Fragment key={task}>
                            <div className="chart-wrapper">
                                <button style={{ float: 'right', marginBottom: 8 }} onClick={() => setFullscreenTask(task)}>
                                    Fullscreen
                                </button>
                                <Line data={chartData} options={options} />
                            </div>
                            {isFullscreen && (
                                <div className="fullscreen-modal" onClick={() => setFullscreenTask(null)}>
                                    <div className="fullscreen-chart" onClick={e => e.stopPropagation()}>
                                        <button style={{ float: 'right', marginBottom: 8 }} onClick={() => setFullscreenTask(null)}>
                                            Close
                                        </button>
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