package ex13_3

import (
	"bufio"
	"bytes"
	"compress/bzip2"
	"fmt"
	"io"
	"strconv"
	"sync"
	"testing"
)

func TestConcurrentWriting(t *testing.T) {
	N := 1000
	c := make(chan int, N)
	for i := 0; i < N; i++ {
		c <- i
	}
	close(c)

	compressed := &bytes.Buffer{}
	writer := NewWriter(compressed)
	var err error
	wg := &sync.WaitGroup{}

	consume := func() {
		defer wg.Done()
		for i := range c {
			_, err = writer.Write([]byte(fmt.Sprintf("%v\n", i)))
			if err != nil {
				return
			}
		}
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go consume()
	}

	wg.Wait()
	if err != nil {
		t.Error(err)
	}
	writer.Close()

	seen := make(map[int]bool)
	uncomprssed := &bytes.Buffer{}
	io.Copy(uncomprssed, bzip2.NewReader(compressed))
	scanner := bufio.NewScanner(uncomprssed)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Error(err)
			return
		}
		seen[i] = true
	}
	var missing []int
	for i := 0; i < N; i++ {
		if !seen[i] {
			missing = append(missing, i)
		}
	}
	if len(missing) > 0 {
		t.Errorf("missing value: %v", missing)
	}
}
