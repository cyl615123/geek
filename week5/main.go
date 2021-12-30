package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"sync/atomic"
	"time"
)

type TimeSlidingWindowCounter struct {
	sync.RWMutex
	sliceNum  int64
	TimeSlice map[int64]int64
}

func NewTimeSlidingWindowCounter(sliceNum int64) *TimeSlidingWindowCounter {
	return &TimeSlidingWindowCounter{
		sliceNum:  sliceNum,
		TimeSlice: make(map[int64]int64),
	}
}

func (sw *TimeSlidingWindowCounter) sliding(now int64) {
	for ts := range sw.TimeSlice {
		if ts+sw.sliceNum < now {
			delete(sw.TimeSlice, ts)
		}
	}
}

func (sw *TimeSlidingWindowCounter) IncrBy(i int64) {
	now := time.Now().Unix()
	sw.Lock()
	defer sw.Unlock()
	sw.TimeSlice[now]++
	sw.sliding(now)
}

func (sw *TimeSlidingWindowCounter) Sum() int64 {
	sum := new(int64)

	sw.RLock()
	defer sw.RUnlock()
	maxOutDateTimeSlice := time.Now().Unix() - sw.sliceNum
	for ts, v := range sw.TimeSlice {
		if ts > maxOutDateTimeSlice {
			atomic.AddInt64(sum, v)
		}
	}
	return *sum
}

func main() {
	counter := NewTimeSlidingWindowCounter(10)
	eg, _ := errgroup.WithContext(context.Background())
	for i := 0; i < 10; i++ {
		eg.Go(func() error {
			timer := time.NewTicker(time.Millisecond)
			for {
				select {
				case <-timer.C:
					counter.IncrBy(1)
				}
			}
		})
	}
	eg.Go(func() error {
		timer := time.NewTicker(time.Second)
		last := int64(0)
		for {
			select {
			case <-timer.C:
				cur := counter.Sum()
				fmt.Printf("cur=%d, diff=%d\n", cur, cur-last)
				last = cur
			}
		}
	})
	eg.Wait()
}
