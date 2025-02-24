package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "proxy_requests_total",
			Help: "Total de solicitudes atendidas por el proxy",
		},
		[]string{"method", "status"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(requestsTotal)
}

func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

func RecordRequest(method, status string) {
	requestsTotal.WithLabelValues(method, status).Inc()
}
