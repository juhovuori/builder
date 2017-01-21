package client

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client is a component that can interact with builder server
type Client interface {
	Shutdown(token string) error
	AddStage(stage string, data io.Reader) error
	Build(buildID string) (string, error)
}

type httpClient struct {
	url     string
	buildID string
}

func (c httpClient) Shutdown(token string) error {
	url := fmt.Sprintf("%s/v1/shutdown", c.url)
	resp, err := http.Post(url, "text/plain", strings.NewReader(token))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Server returned %d", resp.StatusCode)
	}
	return nil
}

func (c httpClient) Build(projectID string) (string, error) {
	url := fmt.Sprintf("%s/v1/projects/%s/trigger", c.url, projectID)
	resp, err := http.Post(url, "text/plain", nil)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Server returned %d", resp.StatusCode)
	}
	build, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(build), nil
}

func (c httpClient) AddStage(stage string, data io.Reader) error {
	url := fmt.Sprintf("%s/v1/build/%s", c.url, c.buildID)
	resp, err := http.Post(url, "text/plain", data)
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

// NewWithoutBuildID creates a new Client
func NewWithoutBuildID(url string) (Client, error) {
	if url == "" {
		return nil, errors.New("Missing builder URL")
	}
	c := httpClient{url, ""}
	return c, nil
}
