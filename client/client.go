package client

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/juhovuori/builder/build"
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
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, string(body))
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
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("%d: %s", resp.StatusCode, string(body))
	}
	build, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(build), nil
}

func (c httpClient) AddStage(name string, data io.Reader) error {
	url := fmt.Sprintf("%s/v1/builds/%s?type=%s&name=%s", c.url, c.buildID, build.PROGRESS, url.QueryEscape(name))
	resp, err := http.Post(url, "text/plain", data)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, string(body))
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
