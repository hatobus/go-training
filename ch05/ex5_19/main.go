package main

import (
	"fmt"
)

func Weird() (str string) {
	defer func() {
		recover()
		str = "hello"
	}()
	panic("panic!")
}

func main() {
	fmt.Println(Weird())
}
