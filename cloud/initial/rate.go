package initial

import (
	"golang.org/x/time/rate"
)

func (initial *Initial) InitRate() *Initial {
	initial.Limiter = rate.NewLimiter(
		rate.Limit(initial.Config.Float("rate.limit", 10)),
		initial.Config.Int("rate.burst", 10),
	)
	return initial
}
