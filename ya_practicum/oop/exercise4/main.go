package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start  time.Time
	splits []time.Duration
}

func (sw *Stopwatch) Start() {
	sw.start = time.Now()
	sw.splits = nil
}

func (sw *Stopwatch) SaveSplit() {
	sw.splits = append(sw.splits, time.Since(sw.start))
}

func (sw Stopwatch) GetResults() []time.Duration {
	return sw.splits
}
func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(2 * time.Second)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
