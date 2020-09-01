package main

import (
	"strings"
)

func EchoJoin(args []string) string {
	return strings.Join(args, " ")
}
