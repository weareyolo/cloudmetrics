package main

import (
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/savaki/cloudmetrics"
)

func main() {
	t := metrics.NewTimer()
	metrics.Register("sample", t)
	t.Update(time.Millisecond)

	// publish metrics to cloudwatch
	cloudmetrics.Publish(nil, metrics.DefaultRegistry, "sample-namespace")
}
