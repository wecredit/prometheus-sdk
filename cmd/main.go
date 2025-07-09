package main

import (
	"log"
	"net/http"
	"github.com/wecredit/prometheus-sdk/config"
	"github.com/wecredit/prometheus-sdk/metrics"
	"github.com/wecredit/prometheus-sdk/server"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	err := config.LoadConfig("sdk_config.json")
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	metrics.Init()

	http.Handle("/metrics", promhttp.Handler())
	routes.RegisterRoutes()

	log.Printf("Prometheus SDK API running on :%s", config.Cfg.MetricsPort)
	log.Fatal(http.ListenAndServe(":"+config.Cfg.MetricsPort, nil))
}
