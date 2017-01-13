package project

import "log"

// Container is the project manager
type Container interface {
	Configure([]string)
	Projects() []Project
	Project(string) (Project, error)
}

type defaultContainer struct {
	projects []Project
}

func (c *defaultContainer) Configure(URLs []string) {
	// TODO: reconfiguration
	for _, URL := range URLs {
		project, err := NewProject(URL)
		if err != nil {
			log.Printf("Could not create project %s: %v\n", URL, err)
		}
		c.projects = append(c.projects, project)
	}
}

func (c *defaultContainer) Projects() []Project {
	return c.projects
}

func (c *defaultContainer) Project(projectID string) (Project, error) {
	for _, pr := range c.projects {
		if pr.ID() == projectID {
			return pr, nil
		}
	}
	return nil, ErrNotFound
}

// NewContainer creates a new project manager
func NewContainer(cfg Config) (Container, error) {
	c := &defaultContainer{}
	c.Configure(cfg.Projects())
	return c, nil
}
