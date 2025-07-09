package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	infoCounter  = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "info_events_total", Help: "Info events"}, []string{"project", "event"})
	errorCounter = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "error_events_total", Help: "Error events"}, []string{"project", "error_type"})
	once         sync.Once
)

func Init() {
	once.Do(func() {
		prometheus.MustRegister(infoCounter)
		prometheus.MustRegister(errorCounter)
	})
}

func IncInfo(project, event string) {
	infoCounter.WithLabelValues(project, event).Inc()
}

func IncError(project, errType string) {
	errorCounter.WithLabelValues(project, errType).Inc()
}
