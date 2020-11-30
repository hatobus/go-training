package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var pool = make(chan struct{}, 20)
var flagV = flag.Bool("v", false, "show progress message")

type DirectorySize struct {
	Root int
	Size int64
}

func directoryEntries(d string) []os.FileInfo {
	//
	pool <- struct{}{}
	defer func() { <-pool }()

	entries, err := ioutil.ReadDir(d)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du error: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(d string, wg *sync.WaitGroup, root int, directorySize chan<- DirectorySize) {
	defer wg.Done()
	for _, entry := range directoryEntries(d) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(d, entry.Name())
			go walkDir(subdir, wg, root, directorySize)
		} else {
			directorySize <- DirectorySize{root, entry.Size()}
		}
	}
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%10d files  %.3f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
}

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{os.Getenv("PWD")}
	}

	dirSizes := make(chan DirectorySize)
	var wg sync.WaitGroup

	for i, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, i, dirSizes)
	}

	go func() {
		wg.Wait()
		close(dirSizes)
	}()

	var tick <-chan time.Time
	if *flagV {
		tick = time.Tick(500 * time.Millisecond)
	}

	fileSum := make([]int64, len(roots))
	byteSize := make([]int64, len(roots))

loop:
	for {
		select {
		case ds, ok := <-dirSizes:
			if !ok {
				break loop
			}
			fileSum[ds.Root]++
			byteSize[ds.Root] += ds.Size
		case <-tick:
			printDiskUsage(roots, fileSum, byteSize)
		}
	}

	printDiskUsage(roots, fileSum, byteSize)
}
