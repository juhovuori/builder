package fetcher

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Fetcher is capable of fetching an URL
type Fetcher interface {
	Fetch(URL string) ([]byte, error)
}

type defaultFetcher struct{}

var (
	// DefaultFetcher is the default fetcher implementation
	DefaultFetcher = defaultFetcher{}
)

func (f defaultFetcher) Fetch(URL string) ([]byte, error) {
	return Fetch(URL)
}

// FetchReadCloser fetches an URL or local file
func FetchReadCloser(URL string) (io.ReadCloser, error) {
	parsed, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	if parsed.Scheme == "" {
		return os.Open(URL)
	}
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// Fetch an URL or local file
func Fetch(URL string) ([]byte, error) {
	data, err := FetchReadCloser(URL)
	if err != nil {
		return nil, err
	}

	defer data.Close()
	return ioutil.ReadAll(data)
}
