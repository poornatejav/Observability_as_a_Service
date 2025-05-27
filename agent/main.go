package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type TelemetryData struct {
	Type    string  `json:"type"`
	Message string  `json:"message"`
	Value   float64 `json:"value"`
	Time    string  `json:"time"`
}

func sendTelemetry(endpoint string) {
	data := TelemetryData{
		Type:    "metric",
		Message: "cpu_usage",
		Value:   rand.Float64() * 100,
		Time:    time.Now().Format(time.RFC3339),
	}

	jsonData, _ := json.Marshal(data)
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending telemetry:", err)
		return
	}
	defer resp.Body.Close()
	log.Println("Sent telemetry:", data)
}

func main() {
	endpoint := os.Getenv("TELEMETRY_ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:8080/ingest"
	}

	rand.Seed(time.Now().UnixNano())
	for {
		sendTelemetry(endpoint)
		time.Sleep(5 * time.Second)
	}
}
