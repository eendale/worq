package job

import (
	"time"
)

// RetryPolicy defines how failed jobs should be retried.
type RetryPolicy struct {
	MaxRetries int
	Delay      time.Duration
}

// ShouldRetry returns whether a job should be retried.
func (r RetryPolicy) ShouldRetry(attempt int) bool {
	return attempt < r.MaxRetries
}

// BackoffDelay returns delay duration before next retry.
func (r RetryPolicy) BackoffDelay(attempt int) time.Duration {
	return time.Duration(attempt) * r.Delay
}
