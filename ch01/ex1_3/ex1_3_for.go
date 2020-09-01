package main

func EchoFor(args []string) string {
	var s, sep string

	for _, arg := range args {
		s += sep + arg
		sep = " "
	}

	return s
}
