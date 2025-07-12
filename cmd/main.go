package main

import (
	"log"
	"net/http"

	"github.com/wecredit/prometheus-sdk/config"
	"github.com/wecredit/prometheus-sdk/metrics"
	routes "github.com/wecredit/prometheus-sdk/server"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Load the configuration from the specified JSON file
	err := config.LoadConfig("sdk_config.json")
	if err != nil {
		log.Fatalf("Config error: %v", err) // Log and exit if there is a configuration error
	}

	// Initialize the metrics system
	metrics.Init()

	// Register the Prometheus metrics handler
	http.Handle("/metrics", promhttp.Handler())

	// Register application-specific routes
	routes.RegisterRoutes()

	// Log the server start message and listen on the configured port
	log.Printf("Prometheus SDK API running on :%s", config.Cfg.MetricsPort)
	log.Fatal(http.ListenAndServe(":"+config.Cfg.MetricsPort, nil)) // Start the HTTP server
}
