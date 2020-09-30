package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	SHA256MODE = "sha256"
	SHA384MODE = "sha384"
	SHA512MODE = "sha512"
)

var (
	HASHTYPE = flag.String("type", "sha256", "hash type")
)

func main() {
	flag.Parse()

	switch *HASHTYPE {
	case SHA256MODE, SHA384MODE, SHA512MODE:
	default:
		fmt.Printf("invalid has type: %v is not supported\n", *HASHTYPE)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}

		bytes := scanner.Bytes()

		switch *HASHTYPE {
		case SHA256MODE:
			hash := sha256.Sum256(bytes)
			fmt.Println(hex.EncodeToString(hash[:]))
		case SHA384MODE:
			hash := sha512.Sum384(bytes)
			fmt.Println(hex.EncodeToString(hash[:]))
		case SHA512MODE:
			hash := sha512.Sum512(bytes)
			fmt.Println(hex.EncodeToString(hash[:]))
		}
	}
}
