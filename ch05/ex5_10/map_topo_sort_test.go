package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMapTopoSort(t *testing.T) {
	// run map_topo_sort.go
	outMapTopo, err := exec.Command("go", "run", "map_topo_sort.go").Output()
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open("./test/golden.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(b), string(outMapTopo)); diff != "" {
		t.Fatalf("invalid output, diff: %v\n", diff)
	}
}
