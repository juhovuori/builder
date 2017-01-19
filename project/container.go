package project

import "log"

// Container is the project manager
type Container interface {
	Projects() []string
	Project(string) (Project, error)
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

// NewContainer creates a new project manager
func NewContainer(projects []string) (Container, error) {
	c := &defaultContainer{map[string]Project{}}
	for _, URL := range projects {
		project, err := NewProject(URL)
		if err != nil {
			log.Printf("Could not create project %s: %v\n", URL, err)
		}
		id := project.ID()
		if _, ok := c.projects[id]; ok {
			log.Printf("Duplicate project %s: %v\n", URL, err)
			// TODO: return error
		}
		c.projects[id] = project
	}
	return c, nil
}
