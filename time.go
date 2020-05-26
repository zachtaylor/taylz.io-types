package types

import "time"

// Hour = time.Hour
var Hour = time.Hour

// Millisecond = time.Millisecond
var Millisecond = time.Millisecond

// Minute = time.Minute
var Minute = time.Minute

// Second = time.Second
var Second = time.Second

// Duration = time.Duration
type Duration = time.Duration

// Ticker = time.Ticker
type Ticker = time.Ticker

// Time = time.Time
type Time = time.Time

// Timer = time.Timer
type Timer = time.Timer

// NewTicker creates a time.Ticker
func NewTicker(d Duration) *Ticker {
	return time.NewTicker(d)
}

// NewTimer creates a time.Timer
func NewTimer(d Duration) *Timer {
	return time.NewTimer(d)
}

// TimeNow returns time.Now()
func TimeNow() Time {
	return time.Now()
}

// TimeUnix return time.Unix()
func TimeUnix(sec int64, nsec int64) Time {
	return time.Unix(sec, nsec)
}
