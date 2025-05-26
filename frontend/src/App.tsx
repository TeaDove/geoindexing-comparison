import { Routes, Route } from 'react-router-dom';
import Chart from './pages/Chart';
import Visualizer from './pages/Visualizer';
import Layout from './components/Layout';
import './App.css'

function App() {
  return (
    <div className="app">
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Chart />} />
          <Route path="charts" element={<Chart />} />
          <Route path="visualizer" element={<Visualizer />} />
        </Route>
      </Routes>
    </div>
  )
}

export default App
