package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	var i int64

	t := time.NewTicker(10 * time.Second)
	defer t.Stop()

	go func() {
		c <- 0
		for {
			i++
			c <- <-c
		}
	}()

	go func() {
		for {
			c <- <-c
		}
	}()

	for {
		select {
		case <-t.C:
			fmt.Printf("%v times per second\n", float64(i/10))
			return
		}
	}
}
