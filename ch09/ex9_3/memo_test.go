package memo

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"
)

func httpGetBody(ctx context.Context, url string) (interface{}, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func TestCancel(t *testing.T) {
	m := New(httpGetBody)

	key := "key"

}
