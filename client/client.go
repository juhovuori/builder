package client

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Client is a component that can interact with builder server
type Client interface {
	Shutdown(token string) error
}

type httpClient struct {
	url     string
	buildID string
}

func (c httpClient) Shutdown(token string) error {
	resp, err := http.Post(c.url+"/v1/shutdown", "text/plain", strings.NewReader(token))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Server returned %d", resp.StatusCode)
	}
	return nil
}

// New creates a new Client
func New(url string, buildID string) (Client, error) {
	if url == "" {
		return nil, errors.New("Missing builder URL")
	}
	if buildID == "" {
		return nil, errors.New("Missing builder buildID")
	}
	c := httpClient{url, buildID}
	return c, nil
}
