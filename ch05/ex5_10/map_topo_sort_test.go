package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMapTopoSort(t *testing.T) {
	u, err := url.Parse("https://raw.githubusercontent.com/adonovan/gopl.io/master/ch5/toposort/main.go")
	if err != nil {
		t.Fatal(err)
	}

	originalGoCodePath := "original.go"

	defer func() {
		if err := os.Remove(originalGoCodePath); !os.IsNotExist(err) || err != nil {
			t.Fatal(err)
		}
	}()

	resp, err := http.Get(u.String())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(originalGoCodePath)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	// run map_topo_sort.go
	outMapTopo, err := exec.Command("go", "run", "map_topo_sort.go").Output()
	if err != nil {
		t.Fatal(err)
	}

	outOrignalTopo, err := exec.Command("go", "run", out.Name()).Output()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(outOrignalTopo, outMapTopo); diff != "" {
		t.Fatalf("invalid output, diff: %v\n", diff)
	}
}
