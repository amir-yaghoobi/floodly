package reporter

import (
	"fmt"
	"io"
	"time"
)

type Standard struct {
	Start  time.Time
	End    time.Time
	Writer io.Writer
}

func (r Standard) Analyze(results <-chan Result) {
	r.Start = time.Now()

	var total int64
	var numberOfErrors int
	var sumOfDurations time.Duration
	var maximumDuration time.Duration
	var minimumDuration = time.Hour * 24

	for result := range results {
		total++
		sumOfDurations += result.Duration

		if result.Error != nil {
			numberOfErrors++
		}

		if maximumDuration < result.Duration {
			maximumDuration = result.Duration
		}
		if minimumDuration > result.Duration {
			minimumDuration = result.Duration
		}

	}

	r.End = time.Now()

	averageDuration := time.Duration(sumOfDurations.Nanoseconds() / total)

	operationDuration := r.End.Sub(r.Start)

	fmt.Fprintf(r.Writer, "process finished after %s\n", operationDuration)
	fmt.Fprintf(r.Writer, "total inserts: %d\n", total)
	fmt.Fprintf(r.Writer, "errors: %d\n", numberOfErrors)
	fmt.Fprintf(r.Writer, "average response time: %s\n", averageDuration)
	fmt.Fprintf(r.Writer, "maximum response time: %s\n", maximumDuration)
	fmt.Fprintf(r.Writer, "minimum response time: %s\n", minimumDuration)
}

func NewStandardReporter(writer io.Writer) *Standard {
	return &Standard{Writer: writer}
}
