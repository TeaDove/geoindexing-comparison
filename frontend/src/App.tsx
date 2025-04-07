import { useState, useEffect } from 'react'
import { Task, Index, RunSettingsType } from './types'
import RunSettings from './components/RunSettings'
import Charts from './components/Charts'
import Notification, { NotificationMessage } from './components/Notification'
import { API_URL } from './config'
import './App.css'

function App() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [indexes, setIndexes] = useState<Index[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [notification, setNotification] = useState<NotificationMessage | null>(null)

  useEffect(() => {
    fetchTasks()
    fetchIndexes()
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
      const response = await fetch(`${API_URL}/tasks`)
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
      const response = await fetch(`${API_URL}/indexes`)
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

  const handleResume = async (settings: RunSettingsType) => {
    setIsLoading(true)
    try {
      const response = await fetch(`${API_URL}/runs/resume`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
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
        headers: {
          'Content-Type': 'application/json',
        },
      })
      showNotification(response.status, '/runs/reset', 'POST')
      if (!response.ok) {
        const error = await response.text()
        showNotification(response.status, '/runs/reset', 'POST', error)
        return
      }
      const data = await response.json()
      console.log('Reset response:', data)
    } catch (error) {
      console.error('Error resetting run:', error)
      showNotification(500, '/runs/reset', 'POST', error instanceof Error ? error.message : 'Unknown error')
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="app">
      <div className="controls">
        <button onClick={() => handleReset()} disabled={isLoading}>
          ⏹
        </button>
      </div>
      <RunSettings
        tasks={tasks}
        indexes={indexes}
        onResume={handleResume}
        isLoading={isLoading}
      />
      <Charts />
      <Notification message={notification} />
    </div>
  )
}

export default App
