package metrics

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	infoCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "info_events_total", Help: "Info events"},
		[]string{"project", "event"},
	)
	errorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "error_events_total", Help: "Error events"},
		[]string{"project", "error_type"},
	)
	infoTimestampGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "info_event_last_timestamp", Help: "Last time an info event occurred (Unix timestamp)"},
		[]string{"project", "event"},
	)
	errorTimestampGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "error_event_last_timestamp", Help: "Last time an error event occurred (Unix timestamp)"},
		[]string{"project", "error_type"},
	)
	once sync.Once
)

func Init() {
	once.Do(func() {
		prometheus.MustRegister(infoCounter)
		prometheus.MustRegister(errorCounter)
		prometheus.MustRegister(infoTimestampGauge)
		prometheus.MustRegister(errorTimestampGauge)
	})
}

func IncInfo(project, event string) {
	infoCounter.WithLabelValues(project, event).Inc()
	infoTimestampGauge.WithLabelValues(project, event).Set(float64(time.Now().Unix()))
}

func IncError(project, errType string) {
	errorCounter.WithLabelValues(project, errType).Inc()
	errorTimestampGauge.WithLabelValues(project, errType).Set(float64(time.Now().Unix()))
}
