package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var cudCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ocp_tip_api_successful_cud_requests_total",
		Help: "Total number of successful CUD requests",
	},
	[]string{"operation"},
)

func RegisterMetrics() {
	prometheus.MustRegister(cudCounter)
}

func IncCudCounter(operation string) {
	cudCounter.With(prometheus.Labels{"operation": operation}).Inc()
}
