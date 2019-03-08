package reporter

import "time"

type Result struct {
	Duration time.Duration
	Error    error
}

type Reporter interface {
	Analyze(results <-chan Result)
}
