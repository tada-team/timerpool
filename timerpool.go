package timerpool

import (
	"sync"
	"time"
)

var pool sync.Pool

func Get(d time.Duration) *time.Timer {
	if v := pool.Get(); v != nil {
		t := v.(*time.Timer)
		if t.Reset(d) {
			panic("timerpool: active timer trapped to the pool")
		}
		return t
	}
	return time.NewTimer(d)
}

func Release(t *time.Timer) {
	if !t.Stop() {
		select {
		case <-t.C:
		default:
		}
	}
	pool.Put(t)
}
