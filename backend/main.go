package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type TelemetryData struct {
	Type    string  `json:"type"`
	Message string  `json:"message"`
	Value   float64 `json:"value"`
	Time    string  `json:"time"`
}

var (
	telemetryStore []TelemetryData
	mu             sync.Mutex
)

func setCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins - adjust if needed
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func ingestHandler(w http.ResponseWriter, r *http.Request) {
	setCors(w)

	if r.Method == http.MethodOptions {
		// Preflight request handling for CORS
		w.WriteHeader(http.StatusOK)
		return
	}

	var data TelemetryData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	telemetryStore = append(telemetryStore, data)
	if len(telemetryStore) > 100 {
		telemetryStore = telemetryStore[len(telemetryStore)-100:]
	}
	mu.Unlock()

	log.Printf("Received metric: %s = %f at %s\n", data.Message, data.Value, data.Time)
	w.WriteHeader(http.StatusOK)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	setCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(telemetryStore); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/ingest", ingestHandler)
	http.HandleFunc("/metrics", metricsHandler)

	log.Println("Starting Ingestion Service on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
