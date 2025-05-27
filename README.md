# Observability MVP 🚀

A minimal but functional **Observability-as-a-Service** prototype built using Go. This MVP showcases the core of a cloud-native monitoring system — capable of ingesting, storing, and visualizing telemetry data in real time.

---
### 1. Clone the Repository

```bash
git clone https://github.com/poornatejav/Observability_as_a_Service.git
cd  Observability_as_a_Service/deploy
```

### 2. Run Docker Compose 

```bash
docker compose up --build
```

### 3. Verify the metrics using curl

```bash
curl http://localhost:8080/metrics
```

### 4. Open dashboard from browser

```bash
http://localhost:3000/
```
