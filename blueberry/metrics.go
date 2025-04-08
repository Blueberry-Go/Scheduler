package blueberry

import "github.com/prometheus/client_golang/prometheus"

const (
	promNamespace = "blueberry" // Prefix for all metrics
	labelTaskName = "task_name"
	labelStatus   = "status"
)

var (
	// --- Built-in Metrics ---
	tasksRegisteredGauge           prometheus.Gauge
	schedulesRegisteredGaugeVec    *prometheus.GaugeVec // GaugeVec to track schedules per task
	tasksExecutingGauge            prometheus.Gauge
	taskExecutionTotalCounter      *prometheus.CounterVec   // Counts executions by task and status
	taskExecutionDurationHistogram *prometheus.HistogramVec // Measures execution duration by task and status
)

// PrometheusRegistry returns the instance's Prometheus registry.
// Used internally and potentially externally (e.g., for the /metrics handler).
func (r *BlueBerry) PrometheusRegistry() *prometheus.Registry {
	return r.promRegistry
}

// MustRegisterMetric allows users to register their own custom metrics
// with this BlueBerry instance's registry. Panics on error.
func (r *BlueBerry) MustRegisterMetric(collector prometheus.Collector) {
	r.promRegistry.MustRegister(collector)
}

// RegisterMetric allows users to register their own custom metrics
// with this BlueBerry instance's registry. Returns error on failure.
func (r *BlueBerry) RegisterMetric(collector prometheus.Collector) error {
	return r.promRegistry.Register(collector)
}
