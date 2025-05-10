import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom';
import type { Task, Index, RunSettings as RunSettingsType, Run } from '../types/index'
import RunSettingsComponent from '../components/RunSettings'
import RunsList from '../components/RunsList'
import Notification, { NotificationMessage } from '../components/Notification'
import { API_URL } from '../config'

const headers = {
    'Content-Type': 'application/json'
};

const Chart: React.FC = () => {
    const [tasks, setTasks] = useState<Task[]>([])
    const [indexes, setIndexes] = useState<Index[]>([])
    const [runs, setRuns] = useState<Run[]>([])
    const [selectedRunId, setSelectedRunId] = useState<number | null>(null)
    const [isLoading, setIsLoading] = useState(false)
    const [notification, setNotification] = useState<NotificationMessage | null>(null)
    const [pointsStart, setPointsStart] = useState(1000)
    const [pointsEnd, setPointsEnd] = useState(10000)
    const [pointsStep, setPointsStep] = useState(100)
    const [displayStart, setDisplayStart] = useState('1 000')
    const [displayEnd, setDisplayEnd] = useState('10 000')
    const [displayStep, setDisplayStep] = useState('100')

    useEffect(() => {
        fetchTasks()
        fetchIndexes()
        fetchRuns()
    }, [])

    const showNotification = (status: number, endpoint: string, method: string, error?: string) => {
        setNotification({
            status,
            endpoint,
            method,
            error,
            timestamp: Date.now(),
        })
    }

    const fetchTasks = async () => {
        try {
            const response = await fetch(`${API_URL}/tasks`, { headers })
            showNotification(response.status, '/tasks', 'GET')
            if (!response.ok) {
                const error = await response.text()
                showNotification(response.status, '/tasks', 'GET', error)
                return
            }
            const data = await response.json()
            setTasks(data)
        } catch (error) {
            console.error('Error fetching tasks:', error)
            showNotification(500, '/tasks', 'GET', error instanceof Error ? error.message : 'Unknown error')
        }
    }

    const fetchIndexes = async () => {
        try {
            const response = await fetch(`${API_URL}/indexes`, { headers })
            showNotification(response.status, '/indexes', 'GET')
            if (!response.ok) {
                const error = await response.text()
                showNotification(response.status, '/indexes', 'GET', error)
                return
            }
            const data = await response.json()
            setIndexes(data)
        } catch (error) {
            console.error('Error fetching indexes:', error)
            showNotification(500, '/indexes', 'GET', error instanceof Error ? error.message : 'Unknown error')
        }
    }

    const fetchRuns = async () => {
        try {
            const response = await fetch(`${API_URL}/runs`, { headers })
            showNotification(response.status, '/runs', 'GET')
            if (!response.ok) {
                const error = await response.text()
                showNotification(response.status, '/runs', 'GET', error)
                return
            }
            const data = await response.json()
            setRuns(data)
        } catch (error) {
            console.error('Error fetching runs:', error)
            showNotification(500, '/runs', 'GET', error instanceof Error ? error.message : 'Unknown error')
        }
    }

    const handleResume = async (settings: RunSettingsType) => {
        setIsLoading(true)
        try {
            const response = await fetch(`${API_URL}/runs/resume`, {
                method: 'POST',
                headers,
                body: JSON.stringify(settings),
            })
            showNotification(response.status, '/runs/resume', 'POST')
            if (!response.ok) {
                const error = await response.text()
                showNotification(response.status, '/runs/resume', 'POST', error)
                return
            }
            const data = await response.json()
            console.log('Resume response:', data)
            await fetchRuns()
        } catch (error) {
            console.error('Error resuming run:', error)
            showNotification(500, '/runs/resume', 'POST', error instanceof Error ? error.message : 'Unknown error')
        } finally {
            setIsLoading(false)
        }
    }

    const handleReset = async () => {
        setIsLoading(true)
        try {
            const response = await fetch(`${API_URL}/runs/reset`, {
                method: 'POST',
                headers,
            })
            showNotification(response.status, '/runs/reset', 'POST')
            if (!response.ok) {
                const error = await response.text()
                showNotification(response.status, '/runs/reset', 'POST', error)
                return
            }
            const data = await response.json()
            console.log('Reset response:', data)
            await fetchRuns()
        } catch (error) {
            console.error('Error resetting run:', error)
            showNotification(500, '/runs/reset', 'POST', error instanceof Error ? error.message : 'Unknown error')
        } finally {
            setIsLoading(false)
        }
    }

    return (
        <div>
            <div className="chart-header-settings">
                <nav>
                    <Link to="/visualizer">Go to Visualizer</Link>
                </nav>
                <div className="points-input" style={{ marginBottom: '1.5rem' }}>
                    <div>
                        <label htmlFor="pointsStart">Начальное кол-во точек</label>
                        <input
                            type="text"
                            id="pointsStart"
                            value={displayStart}
                            onChange={e => {
                                const value = e.target.value;
                                const numericValue = parseInt(value.replace(/\s/g, ''), 10);
                                if (!isNaN(numericValue)) {
                                    setPointsStart(numericValue);
                                    setDisplayStart(numericValue.toLocaleString('ru-RU'));
                                } else {
                                    setDisplayStart(value);
                                }
                            }}
                            min="100"
                            max="1000000"
                        />
                    </div>
                    <div>
                        <label htmlFor="pointsEnd">Конечное кол-во точек</label>
                        <input
                            type="text"
                            id="pointsEnd"
                            value={displayEnd}
                            onChange={e => {
                                const value = e.target.value;
                                const numericValue = parseInt(value.replace(/\s/g, ''), 10);
                                if (!isNaN(numericValue)) {
                                    setPointsEnd(numericValue);
                                    setDisplayEnd(numericValue.toLocaleString('ru-RU'));
                                } else {
                                    setDisplayEnd(value);
                                }
                            }}
                            min="100"
                            max="1000000"
                        />
                    </div>
                    <div>
                        <label htmlFor="pointsStep">Шаг</label>
                        <input
                            type="text"
                            id="pointsStep"
                            value={displayStep}
                            onChange={e => {
                                const value = e.target.value;
                                const numericValue = parseInt(value.replace(/\s/g, ''), 10);
                                if (!isNaN(numericValue)) {
                                    setPointsStep(numericValue);
                                    setDisplayStep(numericValue.toLocaleString('ru-RU'));
                                } else {
                                    setDisplayStep(value);
                                }
                            }}
                            min="1"
                            max="1000"
                        />
                    </div>
                </div>
            </div>
            <div className="chart-layout">
                <aside className="sidebar">
                    <RunSettingsComponent
                        tasks={tasks}
                        indexes={indexes}
                        onResume={handleResume}
                        onReset={handleReset}
                        isLoading={isLoading}
                        pointsStart={pointsStart}
                        pointsEnd={pointsEnd}
                        pointsStep={pointsStep}
                    />
                </aside>
                <main className="main-content">
                    <RunsList
                        runs={runs}
                        onRunSelect={setSelectedRunId}
                        selectedRunId={selectedRunId}
                    />
                </main>
                <Notification message={notification} />
            </div>
        </div>
    )
}

export default Chart; 