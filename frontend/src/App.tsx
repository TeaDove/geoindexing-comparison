import { Routes, Route } from 'react-router-dom';
import Chart from './pages/Chart';
import Visualizer from './pages/Visualizer';
import './App.css'

function App() {
  return (
    <div className="app">
      <Routes>
        <Route path="/" element={<Chart />} />
        <Route path="/visualizer" element={<Visualizer />} />
      </Routes>
    </div>
  )
}

export default App
