package monitoring

import (
	"context"
	"github.com/lyft/flytestdlib/promutils/labeled"

	"github.com/lyft/flytestdlib/contextutils"
	"github.com/lyft/flytestdlib/promutils"
	"github.com/prometheus/client_golang/prometheus"
)

// Represents a Gauge labeled with values from the context. See labeled.SetMetricsKeys for information about to
// configure that.
type Gauge struct {
	*prometheus.GaugeVec

	prometheus.Gauge
}

// Inc increments the Gauge by 1. Use Add to increment it by arbitrary non-negative values. The data point will be
// labeled with values from context. See labeled.SetMetricsKeys for information about to configure that.
func (g Gauge) Inc(ctx context.Context) {
	gauge, err := g.GaugeVec.GetMetricWith(contextutils.Values(ctx, metricKeys...))
	if err != nil {
		panic(err.Error())
	}
	gauge.Inc()

	if g.Gauge != nil {
		g.Gauge.Inc()
	}
}

// Dec decrements the Gauge by 1. The data point will be labeled with values from context.
// See labeled.SetMetricsKeys for information about to configure that.
func (g Gauge) Dec(ctx context.Context) {
	gauge, err := g.GaugeVec.GetMetricWith(contextutils.Values(ctx, metricKeys...))
	if err != nil {
		panic(err.Error())
	}
	gauge.Dec()

	if g.Gauge != nil {
		g.Gauge.Dec()
	}
}

// Set set the given value to the gauge. It panics if the value is < 0.. The data point will be labeled with values
// from context. See labeled.SetMetricsKeys for information about to configure that.
func (g Gauge) Set(ctx context.Context, v float64) {
	gauge, err := g.GaugeVec.GetMetricWith(contextutils.Values(ctx, metricKeys...))
	if err != nil {
		panic(err.Error())
	}
	gauge.Set(v)

	if g.Gauge != nil {
		g.Gauge.Set(v)
	}
}

// Creates a new labeled gauge. Label keys must be set before instantiating a gauge. See labeled.SetMetricsKeys for
// information about to configure that.
func NewGauge(name, description string, scope promutils.Scope, opts ...labeled.MetricOption) Gauge {
	if len(metricKeys) == 0 {
		panic(ErrNeverSet)
	}

	gauge := Gauge{
		GaugeVec: scope.MustNewGaugeVec(name, description, metricStringKeys...),
	}

	for _, opt := range opts {
		if _, emitUnlabeledMetric := opt.(labeled.EmitUnlabeledMetricOption); emitUnlabeledMetric {
			gauge.Gauge = scope.MustNewGauge(labeled.GetUnlabeledMetricName(name), description)
		}
	}

	return gauge
}
