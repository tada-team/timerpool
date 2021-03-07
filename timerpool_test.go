package timerpool

import (
	"context"
	"sync"
	"testing"
	"time"
)

func BenchmarkTimer(b *testing.B) {
	b.Run("cancelled", func(b *testing.B) {
		dur := 10 * time.Second

		b.Run("newTimer", func(b *testing.B) {
			var wg sync.WaitGroup
			defer wg.Wait()

			ctx, cancel := context.WithCancel(context.Background())

			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					timer := time.NewTimer(dur)
					defer timer.Stop()

					select {
					case <-timer.C:
						b.Fatal("timer executed")
					case <-ctx.Done():
						return
					}
				}()
			}

			cancel()
		})

		b.Run("poolTimer", func(b *testing.B) {
			var wg sync.WaitGroup
			defer wg.Wait()

			ctx, cancel := context.WithCancel(context.Background())

			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					timer := Get(dur)
					defer Release(timer)

					select {
					case <-timer.C:
						b.Fatal("timer executed")
					case <-ctx.Done():
						return
					}
				}()
			}

			cancel()
		})
	})

	b.Run("used", func(b *testing.B) {
		dur := 100 * time.Millisecond

		b.Run("newTimer", func(b *testing.B) {
			var wg sync.WaitGroup
			defer wg.Wait()

			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					timer := time.NewTimer(dur)
					defer timer.Stop()
					<-timer.C
				}()
			}
		})

		b.Run("poolTimer", func(b *testing.B) {
			var wg sync.WaitGroup
			defer wg.Wait()

			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					timer := Get(dur)
					defer Release(timer)

					<-timer.C
				}()
			}
		})
	})
}
