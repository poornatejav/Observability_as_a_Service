import { useEffect, useState } from 'react';
import {
  LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer,
} from 'recharts';

function App() {
  const [metrics, setMetrics] = useState([]);

  useEffect(() => {
    async function fetchMetrics() {
      try {
        const res = await fetch('http://localhost:8080/metrics'); // use correct URL or proxy
        const data = await res.json();
        setMetrics(data);
      } catch (err) {
        console.error('Error fetching metrics:', err);
      }
    }

    fetchMetrics();
    const interval = setInterval(fetchMetrics, 5000);
    return () => clearInterval(interval);
  }, []);

  // Filter only cpu_usage metrics (in case you add more types later)
  const cpuMetrics = metrics.filter(m => m.message === 'cpu_usage');

  // Format timestamps for X-axis (HH:mm:ss)
  const formattedData = cpuMetrics.map(m => ({
    ...m,
    timeFormatted: new Date(m.time).toLocaleTimeString(),
  }));

  const latestValue = cpuMetrics.length > 0 ? cpuMetrics[cpuMetrics.length - 1].value.toFixed(2) : 'N/A';

  return (
    <div style={{ padding: 20, maxWidth: 800, margin: '0 auto' }}>
      <h1>Telemetry Dashboard</h1>

      {cpuMetrics.length === 0 ? (
        <p>No metrics received yet</p>
      ) : (
        <>
          <ResponsiveContainer width="100%" height={300}>
            <LineChart data={formattedData} margin={{ top: 20, right: 30, left: 20, bottom: 5 }}>
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis dataKey="timeFormatted" />
              <YAxis domain={[0, 'dataMax + 10']} />
              <Tooltip />
              <Line type="monotone" dataKey="value" stroke="#8884d8" dot={false} />
            </LineChart>
          </ResponsiveContainer>
          <p>
            Latest CPU Usage: <strong>{latestValue}%</strong>
          </p>
        </>
      )}
    </div>
  );
}

export default App;
