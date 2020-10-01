package command

import (
	"fmt"
	"log"
	"os"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.12/comic"
)

func GenerateIndex(fname string) error {
	comicchan, err := comic.DownloadComics()
	if err != nil {
		return err
	}

	indexNumber := comic.GenerateIndex(comicchan)

	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Println(indexNumber)

	for key, val := range indexNumber {
		fmt.Fprintf(f, fmt.Sprintf("%v: %v\n", key, val))
	}

	return nil
}
