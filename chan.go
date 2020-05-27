package types

import "time"

// NewChanAfter returns time.After()
func NewChanAfter(d Duration) <-chan Time { return time.After(d) }

// NewChanTick returns time.Tick()
func NewChanTick(d Duration) <-chan Time { return time.Tick(d) }
