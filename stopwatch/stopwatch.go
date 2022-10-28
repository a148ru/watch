package stopwatch

import (
	"time"
)

type Stopwatch struct {
	startTime time.Time
	splits    []time.Time
}

func (sw *Stopwatch) Start() {
	sw.startTime = time.Now()
	sw.splits = nil
}

func (sw *Stopwatch) SaveSplit() {
	sw.splits = append(sw.splits, time.Now())
}

func (sw Stopwatch) GetResults() (retResults []time.Duration) {
	for _, splitTime := range sw.splits {
		retResults = append(retResults, splitTime.Sub(sw.startTime))
	}
	return
}
