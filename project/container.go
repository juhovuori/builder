package project

// Container is the project manager
type Container interface {
	Configure([]string)
	Projects() []Project
	Project(string) (Project, error)
}

type container struct {
	projects []Project
}

func (c *container) Configure(URLs []string) {
	// TODO: reconfiguration
	for _, URL := range URLs {
		project := newProject(URL)
		c.projects = append(c.projects, project)
	}
}

func (c *container) Projects() []Project {
	return c.projects
}

func (c *container) Project(projectID string) (Project, error) {
	for _, pr := range c.projects {
		if pr.ID() == projectID {
			return pr, nil
		}
	}
	return nil, ErrNotFound
}

// NewContainer creates a new project manager
func NewContainer(cfg ProjectsConfig) (Container, error) {
	c := &container{}
	c.Configure(cfg.Projects())
	return c, nil
}
