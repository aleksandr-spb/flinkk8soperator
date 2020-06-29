package monitoring

import (
	"context"
	"testing"

	"github.com/lyft/flytestdlib/contextutils"
	"github.com/lyft/flytestdlib/promutils"
	"github.com/stretchr/testify/assert"
)

func TestLabeledGauge(t *testing.T) {
	assert.NotPanics(t, func() {
		SetMetricKeys(contextutils.ProjectKey, contextutils.DomainKey, contextutils.WorkflowIDKey, contextutils.TaskIDKey)
	})

	scope := promutils.NewTestScope()
	gauge := NewGauge("lbl_gauge", "help", scope)
	assert.NotNil(t, gauge)
	ctx := context.TODO()
	gauge.Inc(ctx)
	gauge.Set(ctx, 1.0)
	gauge.Dec(ctx)

	ctx = contextutils.WithProjectDomain(ctx, "project", "domain")
	gauge.Inc(ctx)
	gauge.Set(ctx, 1.0)
	gauge.Dec(ctx)

	ctx = contextutils.WithTaskID(ctx, "task")
	gauge.Inc(ctx)
	gauge.Set(ctx, 1.0)
	gauge.Dec(ctx)
}
