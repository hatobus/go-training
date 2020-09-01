package pioutil

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
)

func OutputCapture(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	stdout := os.Stdout
	stderr := os.Stderr

	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer

	out := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	wg.Wait()
	f()

	err = writer.Close()
	if err != nil {
		log.Printf("writer close failed: %v", err)
	}

	return <- out
}
