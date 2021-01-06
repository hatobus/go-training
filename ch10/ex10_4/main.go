package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func packages(packageName []string) []string {
	args := []string{"list", "-f", "'{{.ImportPath}}'"}
	args = append(args, packageName...)

	out, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Println("packages resolve error:", err)
	}
	return strings.Fields(string(out))
}

func getPkgAncestors(pkgNames []string) []string {
	dest := make(map[string]bool)
	for _, pkg := range pkgNames {
		dest[pkg] = true
	}

	args := []string{"list", "-f", `'{{.ImportPath}} {{join .Deps " "}}'`, "..."}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Println("go list ... failed:", err)
		return nil
	}
	var pkgs []string
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		pkg := fields[0]
		dependency := fields[1:]
		for _, dep := range dependency {
			if dest[dep] {
				pkgs = append(pkgs, pkg)
				break
			}
		}
	}
	return pkgs
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("invalid arguments, ")
	}
	pkgs := getPkgAncestors(packages(os.Args[1:]))
	sort.StringSlice(pkgs).Sort()
	for _, pkg := range pkgs {
		fmt.Println(pkg)
	}
}
