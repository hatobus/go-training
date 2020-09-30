package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	filename = flag.CommandLine.String("file", "", "filename")
)

func main() {
	flag.Parse()

	if *filename == "" {
		log.Fatalf("invalid filename: %v not found", *filename)
	}

	absPath, err := filepath.Abs(*filename)
	if err != nil {
		log.Fatal(err)
	}

	fp, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	wordFrequency := make(map[string]int)

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		w := scanner.Text()
		wordFrequency[w]++
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Printf("%v\n", wordFrequency)
}
