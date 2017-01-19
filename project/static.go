package project

type staticProject struct {
	defaultProject
	id string
}

func (p *staticProject) Description() string { return "" }
func (p *staticProject) Script() string      { return "" }
func (p *staticProject) URL() string         { return "" }
func (p *staticProject) ID() string          { return p.id }
func (p *staticProject) Err() error          { return nil }

// NewStaticProject returns a new project based on static data
func NewStaticProject(id string) Project {
	p := staticProject{
		id: id,
	}
	return &p
}

// NewStaticContainer returns a container with a project for testing
func NewStaticContainer(p Project) Container {
	return &defaultContainer{
		map[string]Project{
			p.ID(): p,
		},
	}
}
