package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const (
	baseURL = "http://localhost:2112"
	project = "nurture-engine"
)

type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	infoEvents  = []string{"init_complete", "job_dispatched", "cache_updated", "heartbeat", "sdk_start"}
	errorEvents = []string{"db_timeout", "cache_miss", "sqs_failure", "panic_recovered", "api_down"}
)

func main() {
	fmt.Println("Testing Prometheus SDK Locally with Simulated Logs")

	healthCheck()

	for i := 1; i <= 5; i++ {
		// Random info event
		info := infoEvents[rand.Intn(len(infoEvents))]
		sendInfo(info)

		// Random error event
		errType := errorEvents[rand.Intn(len(errorEvents))]
		sendError(errType)

		time.Sleep(1 * time.Second)
	}

	// View raw /metrics output (just first few lines)
	printMetrics()

	fmt.Println("Done. Now visit http://localhost:9090 to query.")
	fmt.Println("\nExample Prometheus queries:")
	fmt.Printf("Total info for project:\n    info_events_total{project=\"%s\"}\n", project)
	fmt.Printf("Info for specific event:\n    info_events_total{project=\"%s\",event=\"job_dispatched\"}\n", project)
	fmt.Printf("Errors by type:\n    error_events_total{project=\"%s\",error_type=\"db_timeout\"}\n", project)
}

func healthCheck() {
	fmt.Println("/health check:")
	resp, err := http.Get(baseURL + "/health")
	handleResponse(resp, err)
}

func sendInfo(event string) {
	fmt.Printf("Logging info event: %s\n", event)
	url := fmt.Sprintf("%s/info?project=%s&event=%s", baseURL, project, event)
	resp, err := http.Get(url)
	handleResponse(resp, err)
}

func sendError(errorType string) {
	fmt.Printf("Logging error event: %s\n", errorType)
	url := fmt.Sprintf("%s/error?project=%s&type=%s", baseURL, project, errorType)
	resp, err := http.Get(url)
	handleResponse(resp, err)
}

func printMetrics() {
	fmt.Println("\n/metrics snapshot:")
	resp, err := http.Get(baseURL + "/metrics")
	if err != nil {
		fmt.Println("Could not fetch metrics:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	lines := string(body)
	fmt.Println(lines[:min(len(lines), 1000)]) // limit output
}

func handleResponse(resp *http.Response, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var data ApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Invalid response:", string(body))
		return
	}
	fmt.Printf("%s: %s\n\n", data.Status, data.Message)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
