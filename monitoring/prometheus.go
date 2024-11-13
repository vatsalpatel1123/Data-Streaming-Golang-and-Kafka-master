package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	RequestsProcessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_processed_total",
			Help: "Total number of processed requests",
		},
		[]string{"endpoint"},
	)
	ErrorsOccurred = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "errors_occurred_total",
			Help: "Total number of errors occurred",
		},
		[]string{"endpoint"},
	)
)

func InitPrometheus() {
	prometheus.MustRegister(RequestsProcessed)
	prometheus.MustRegister(ErrorsOccurred)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil)
}
