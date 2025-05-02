import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom';
import { Task, Index, RunSettingsType, Run } from '../types'
import RunSettings from '../components/RunSettings'
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
        <div className="page-container">
            <nav>
                <Link to="/visualizer">Go to Visualizer</Link>
            </nav>
            <RunSettings
                tasks={tasks}
                indexes={indexes}
                onResume={handleResume}
                onReset={handleReset}
                isLoading={isLoading}
            />
            <RunsList
                runs={runs}
                onRunSelect={setSelectedRunId}
                selectedRunId={selectedRunId}
            />
            <Notification message={notification} />
        </div>
    )
}

export default Chart; 