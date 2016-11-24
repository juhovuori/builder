package projects

import "github.com/juhovuori/builder/cfg"

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

func (p *projects) Project(project string) (Project, error) {
	return nil, nil
}

// New creates a new project manager
func New(c cfg.Cfg) (Projects, error) {
	p := &projects{}
	p.Configure(c.Projects())
	return p, nil
}
