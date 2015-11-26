package tick

import (
	"time"
)

// A Ticker holds a channel that delivers `ticks' of a clock at intervals.
type Ticker struct {
	C    <-chan time.Time // The channel on which the ticks are delivered.
	quit chan struct{}
}

// NewTicker returns a new Ticker containing a channel that will send the
// time with a period specified by the duration argument. Times are rounded
// to the period, and no ticks are dropped. Ticks are not guaranteed to be
// on time.
func NewTicker(d time.Duration) *Ticker {
	var (
		c = make(chan time.Time)
		t = &Ticker{
			quit: make(chan struct{}),
			C:    c,
		}
	)
	go t.run(d, c)
	return t
}

// Stop turns off a ticker. After Stop, no more ticks will be sent.
// Stop does not close the channel, to prevent a read from the channel
// succeeding incorrectly.
func (t *Ticker) Stop() {
	close(t.quit)
}

func (t *Ticker) run(d time.Duration, c chan<- time.Time) {
	var (
		start = time.Now()
		first = start.Add(d / 2).Round(d)
	)
	time.Sleep(first.Sub(start))
	ticker := time.NewTicker(d)
	c <- first
	next := first.Add(d)
	for {
		select {
		case x := <-ticker.C:
			for ; next.Before(x); next = next.Add(d) {
				c <- next
			}
		case <-t.quit:
			ticker.Stop()
			return
		}
	}
}

// Tick is a convenience wrapper for NewTicker providing access to the
// ticking channel only.
func Tick(d time.Duration) <-chan time.Time {
	return NewTicker(d).C
}
