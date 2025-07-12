package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// infoCounter tracks the total number of informational events, labeled by project and event type.
	infoCounter = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "info_events_total", Help: "Info events"}, []string{"project", "event"})

	// errorCounter tracks the total number of error events, labeled by project and error type.
	errorCounter = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "error_events_total", Help: "Error events"}, []string{"project", "error_type"})

	// once ensures that the initialization code runs only once.
	once sync.Once
)

// Init registers the Prometheus metrics with the default registry. This function is safe to call multiple times.
func Init() {
	once.Do(func() {
		prometheus.MustRegister(infoCounter)
		prometheus.MustRegister(errorCounter)
	})
}

// IncInfo increments the infoCounter for the given project and event type.
func IncInfo(project, event string) {
	infoCounter.WithLabelValues(project, event).Inc()
}

// IncError increments the errorCounter for the given project and error type.
func IncError(project, errType string) {
	errorCounter.WithLabelValues(project, errType).Inc()
}
