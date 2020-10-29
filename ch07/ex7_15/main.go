package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	eval "github.com/hatobus/go-training/ch07/ex7_14"
)

const (
	// exit status code
	transformerror = 1
	assignError    = 2
)

func main() {
	var exitCode int
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please input: ")

	stdin.Scan()

	exprString := stdin.Text()

	fmt.Printf("<var>=<value>, (ex: a=1) ")

	stdin.Scan()

	envString := stdin.Text()

	if stdin.Err() != nil {
		fmt.Fprintln(os.Stderr, stdin.Err())
		os.Exit(1)
	}

	env := eval.Env{}

	assignments := strings.Fields(envString)

	for _, assign := range assignments {
		fields := strings.Split(assign, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad assignment: %s \n", assign)
			exitCode = assignError
		}
		variable, value := fields[0], fields[1]
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid value %v, using 0 instead of your input. err = %v", variable, err)
			exitCode = transformerror
		}
		env[eval.Var(variable)] = v
	}

	expr, err := eval.Parse(exprString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad expression: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(expr.Eval(env))

	os.Exit(exitCode)
}
