# timerpool

Usage:
```go
import (
	"time"
	"github.com/tada-team/timerpool"
)

func example() {
    dur := time.Second
    
    // without pool  
    timer := time.NewTimer(dur)
    defer timer.Stop()
    
    // with pool
    timer := timerpool.Get(dur)
    defer timerpool.Release(timer)
}
```

```text
BenchmarkTimer/cancelled/newTimer-12         	  343038	      3464 ns/op	     846 B/op	       5 allocs/op
BenchmarkTimer/cancelled/poolTimer-12        	 1000000	      2669 ns/op	     729 B/op	       5 allocs/op
BenchmarkTimer/used/newTimer-12              	 2613874	       516.9 ns/op	     210 B/op	       3 allocs/op
BenchmarkTimer/used/poolTimer-12             	 2319859	       458.1 ns/op	      13 B/op	       0 allocs/op
```
