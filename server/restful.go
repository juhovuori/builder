package server

import (
	"fmt"
	"strconv"

	"github.com/juhovuori/builder/build"
	"github.com/juhovuori/builder/project"
	"github.com/labstack/echo"
)

// Root response
type Root struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Projects string `json:"projects"`
	Builds   string `json:"builds"`
}

// Items response
type Items struct {
	Items []string `json:"items"`
	More  string   `json:"more,omitempty"`
}

var root = Root{
	"Builder",
	"/v1/version",
	"/v1/projects",
	"/v1/builds",
}

// Build response
type Build struct {
	Project      string  `json:"project"`
	Script       string  `json:"script"`
	ExecutorType string  `json:"executor-type"`
	Created      int64   `json:"created"`
	Stages       []Stage `json:"stages"`
	Stdout       string  `json:"stdout"`
}

// Stage response
type Stage struct {
	Type      string `json:"type"`
	Timestamp int64  `json:"timestamp"`
	Name      string `json:"name"`
	Data      string `json:"data"`
}

// Project response
type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Script      string `json:"script"`
	Config      string `json:"config"`
	URL         string `json:"url"`
	VCS         string `json:"vcs"`
	Builds      string `json:"builds"`
	Err         error  `json:"error,omitempty"`
}

func restfulProject(p project.Project) Project {
	return Project{
		Name:        p.Name(),
		Description: p.Description(),
		Script:      p.Script(),
		Config:      p.Config(),
		URL:         p.URL(),
		VCS:         p.VCS(),
		Builds:      fmt.Sprintf("/v1/projects/%s/builds", p.ID()),
		Err:         p.Err(),
	}
}

func restfulBuild(b build.Build) Build {
	rb := Build{
		Project:      fmt.Sprintf("/v1/projects/%s", b.ProjectID()),
		Script:       b.Script(),
		ExecutorType: b.ExecutorType(),
		Created:      b.Created(),
		Stages:       []Stage{},
		Stdout:       fmt.Sprintf("/v1/builds/%s/stdout", b.ID()),
	}
	for i, s := range b.Stages() {
		rb.Stages = append(rb.Stages, Stage{
			Type:      string(s.Type),
			Timestamp: s.Timestamp,
			Name:      s.Name,
			Data:      fmt.Sprintf("/v1/builds/%s/data/%d", b.ID(), i),
		})
	}
	return rb
}

// restfulItems response
func restfulItems(items []string, baseURL string, index, page int) Items {
	res := Items{[]string{}, ""}
	if index > len(items) {
		index = len(items)
	}
	for i, b := range items[index:] {
		if i >= page {
			res.More = fmt.Sprintf("%s?index=%d&page=%d", baseURL, index+page, page)
			break
		}
		res.Items = append(res.Items, fmt.Sprintf("%s/%s", baseURL, b))
	}
	return res
}

func paging(c echo.Context) (int, int, error) {
	indexS := c.QueryParam("index")
	if indexS == "" {
		indexS = "0"
	}
	index, err := strconv.Atoi(indexS)
	if err != nil {
		return 0, 0, err
	}
	pageS := c.QueryParam("page")
	if pageS == "" {
		pageS = "50"
	}
	page, err := strconv.Atoi(pageS)
	if err != nil {
		return 0, 0, err
	}
	return index, page, nil
}
