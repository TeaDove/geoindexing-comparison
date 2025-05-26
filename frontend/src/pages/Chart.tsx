import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom';
import type { Task, Index, RunSettings as RunSettingsType, Run } from '../types/index'
import RunSettingsComponent from '../components/RunSettings'
import RunsList from '../components/RunsList'
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

    const fetchTasks = async () => {
        try {
            const response = await fetch(`${API_URL}/tasks`, { headers })
            if (!response.ok) {
                const error = await response.text()
                console.error('Error fetching tasks:', error)
                return
            }
            const data = await response.json()
            setTasks(data)
        } catch (error) {
            console.error('Error fetching tasks:', error)
        }
    }

    const fetchIndexes = async () => {
        try {
            const response = await fetch(`${API_URL}/indexes`, { headers })
            if (!response.ok) {
                const error = await response.text()
                console.error('Error fetching indexes:', error)
                return
            }
            const data = await response.json()
            setIndexes(data)
        } catch (error) {
            console.error('Error fetching indexes:', error)
        }
    }

    const fetchRuns = async () => {
        try {
            const response = await fetch(`${API_URL}/runs`, { headers })
            if (!response.ok) {
                const error = await response.text()
                console.error('Error fetching runs:', error)
                return
            }
            const data = await response.json()
            setRuns(data)
        } catch (error) {
            console.error('Error fetching runs:', error)
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
            if (!response.ok) {
                const error = await response.text()
                console.error('Error resuming run:', error)
                return
            }
            const data = await response.json()
            console.log('Resume response:', data)
            await fetchRuns()
        } catch (error) {
            console.error('Error resuming run:', error)
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
            if (!response.ok) {
                const error = await response.text()
                console.error('Error resetting run:', error)
                return
            }
            const data = await response.json()
            console.log('Reset response:', data)
            await fetchRuns()
        } catch (error) {
            console.error('Error resetting run:', error)
        } finally {
            setIsLoading(false)
        }
    }

    return (
        <div className="page-container">
            <nav>
                <Link to="/visualizer">Go to Visualizer</Link>
            </nav>
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
            </div>
        </div>
    )
}

export default Chart 