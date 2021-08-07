package static

import "time"

type config struct {
	spawnAfter time.Duration
}

type option func(*config)

func WithDuration(d time.Duration) option {
	return func(c *config) {
		c.spawnAfter = d
	}
}
