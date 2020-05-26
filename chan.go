package types

import "time"

// ChanAfter returns time.After()
func ChanAfter(d Duration) <-chan Time {
	return time.After(d)
}

// ChanTick uses time.Tick
func ChanTick(d Duration) <-chan Time { return time.Tick(d) }
