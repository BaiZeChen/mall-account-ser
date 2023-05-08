package pkg

import (
	"golang.org/x/time/rate"
	"time"
)

var GlobalLimiter *rate.Limiter

func FlowControl() {
	GlobalLimiter = rate.NewLimiter(rate.Every(2*time.Second/10), 1000)
}
