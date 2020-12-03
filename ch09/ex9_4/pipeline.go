package main

import (
	"fmt"
	"time"
)

func pipeline(num int) (chan struct{}, chan struct{}) {
	in := make(chan struct{})
	out := make(chan struct{})
	first := out

	for i := 0; i < num; i++ {
		in = out
		out = make(chan struct{})
		go func(in, out chan struct{}) {
			for val := range in {
				out <- val
			}
			close(out)
		}(in, out)
	}

	return first, out
}

func main() {
	cnt := 1
	for {
		in, out := pipeline(cnt)

		start := time.Now()

		in <- struct{}{}
		<-out

		fmt.Printf("number of pipelines: %v, duration: %v\n", cnt, time.Now().Sub(start))
		cnt++
		close(in)
	}
}
