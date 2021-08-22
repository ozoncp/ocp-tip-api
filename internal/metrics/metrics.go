package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var cudCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "tips_cud_operations_count",
		Help: "Count of CUD operations for tips",
	},
	[]string{"operation"},
)

func RegisterMetrics() {
	prometheus.MustRegister(cudCounter)
}

func IncCudCounter(operation string) {
	cudCounter.With(prometheus.Labels{"operation": operation}).Inc()
}
