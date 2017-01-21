package project

// Container is the project manager
type Container interface {
	Projects() []string
	Project(string) (Project, error)
	Add(project Project) error
}

type defaultContainer struct {
	projects map[string]Project
}

func (c *defaultContainer) Projects() []string {
	projects := []string{}
	for ID := range c.projects {
		projects = append(projects, ID)
	}
	return projects
}

func (c *defaultContainer) Project(projectID string) (Project, error) {
	p, ok := c.projects[projectID]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (c *defaultContainer) Add(project Project) error {
	id := project.ID()
	if _, ok := c.projects[id]; ok {
		return ErrDuplicate
	}
	c.projects[id] = project
	return nil
}

// NewContainer creates a new project manager
func NewContainer() Container {
	c := &defaultContainer{map[string]Project{}}
	return c
}
