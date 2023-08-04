package metrics

import (
	"time"
)

// ITimer is the interface that emits timer metrics.
type ITimer interface {

	// Record records the duration since timer.start, and reset timer.start.
	Record() time.Duration

	// RecordDuration records duration into timer, and reset timer.start.
	RecordDuration(duration time.Duration)

	// Reset resets the timer.start.
	Reset()
}

// timer defines the timer metric. timer is reported to each external Sink-able system.
type timer struct {
	name  string
	start time.Time
}

// Record records the time cost since last Record.
func (t *timer) Record() time.Duration {
	duration := time.Since(t.start)
	t.start = time.Now()
	r := NewSingleDimensionMetrics(t.name, float64(duration), PolicyTimer)
	for _, sink := range metricsSinks {
		sink.Report(r)
	}
	return duration
}

// RecordDuration records duration and reset t.start to now.
func (t *timer) RecordDuration(duration time.Duration) {
	t.start = time.Now()
	r := NewSingleDimensionMetrics(t.name, float64(duration), PolicyTimer)
	for _, sink := range metricsSinks {
		sink.Report(r)
	}
}

// Reset resets the start time of timer to now.
func (t *timer) Reset() {
	t.start = time.Now()
}
