import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [telemetry, setTelemetry] = useState([]);

  useEffect(() => {
    const interval = setInterval(() => {
      fetch("http://localhost:9090/api/v1/query?query=cpu_usage")
        .then(res => res.json())
        .then(data => setTelemetry(data.data.result || []));
    }, 5000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="App">
      <h1>Observability Dashboard</h1>
      <ul>
        {telemetry.map((item, index) => (
          <li key={index}>{JSON.stringify(item)}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;
