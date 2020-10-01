package comic

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"golang.org/x/xerrors"
)

func GetNumberOfComics() (int, error) {
	resp, err := http.Get(URL + "/info.0.json")
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return -1, xerrors.Errorf("failed to get number of comics")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	var c Comic
	if err = json.Unmarshal(b, &c); err != nil {
		return -1, err
	}

	return c.Num, nil
}

func GetComic(n int) (*Comic, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, strconv.Itoa(n), "info.0.json")

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var c Comic
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func comicGetter(number chan int, c chan Comic, done chan int) {
	for n := range number {
		cm, err := GetComic(n)
		if err != nil {
			log.Printf("faild to get comic no %v, reason: %v", n, err)
			continue
		}
		c <- *cm
	}
	done <- 1
}

func DownloadComics() (chan Comic, error) {
	latestNum, err := GetNumberOfComics()
	if err != nil {
		return nil, err
	}

	latestNum = latestNum / 100

	log.Println(latestNum)

	workers := 10
	comics := make(chan Comic, 5*workers)
	sumOfComics := make(chan int, 1*workers)
	done := make(chan int, 0)

	for i := 0; i < workers; i++ {
		go comicGetter(sumOfComics, comics, done)
	}

	for i := 1; i <= latestNum; i++ {
		sumOfComics <- i
	}
	close(sumOfComics)

	for i := 0; i < workers; i++ {
		<-done
	}

	close(done)
	close(comics)
	return comics, nil
}
