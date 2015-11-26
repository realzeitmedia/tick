package tick

import (
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	now := time.Now()
	tick := NewTicker(1 * time.Millisecond)

	t0 := <-tick.C

	if !t0.After(now) {
		t.Errorf("first tick before start: now=%q t0=%q", now, t0)
	}

	for i := 1; i < 10; i++ {
		ti := <-tick.C
		if have, want := ti.Sub(t0), time.Duration(i)*time.Millisecond; have != want {
			t.Errorf("bad %dth tick: have %q, want %q", i, have, want)
		}
	}

	tick.Stop()
}
