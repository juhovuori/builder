package project

// Projects is the project manager
type Projects interface {
	Configure([]string)
	Projects() []Project
	Project(string) (Project, error)
}

type projects struct {
	projects []Project
}

func (p *projects) Configure(URLs []string) {
	// TODO: reconfiguration
	for _, URL := range URLs {
		project := newProject(URL)
		p.projects = append(p.projects, project)
	}
}

func (p *projects) Projects() []Project {
	return p.projects
}

func (p *projects) Project(projectID string) (Project, error) {
	for _, pr := range p.projects {
		if pr.ID() == projectID {
			return pr, nil
		}
	}
	return nil, ErrNotFound
}

// New creates a new project manager
func New(c ProjectsConfig) (Projects, error) {
	p := &projects{}
	p.Configure(c.Projects())
	return p, nil
}
