package build

import "github.com/juhovuori/builder/project"

// Build describes a single build
type Build interface {
	ID() string
	ExecutorType() string
	ProjectID() string
	Script() string
	Completed() bool
	Error() error
}

type defaultBuild struct {
	id           string
	projectID    string
	script       string
	executorType string
	completed    bool
	err          error
}

func (b *defaultBuild) ID() string {
	return b.id
}

func (b *defaultBuild) ExecutorType() string {
	return b.executorType
}

func (b *defaultBuild) ProjectID() string {
	return b.projectID
}

func (b *defaultBuild) Script() string {
	return b.script
}

func (b *defaultBuild) Completed() bool {
	return b.completed
}

func (b *defaultBuild) Error() error {
	return b.err
}

// New creates a new build
func New(project project.Project) (Build, error) {
	if project == nil {
		return nil, ErrNilProject
	}
	e := "nop"
	if project.Script() != "" {
		e = "fork"
	}
	b := defaultBuild{
		id:           "",
		projectID:    project.ID(),
		script:       project.Script(),
		executorType: e,
		completed:    false,
		err:          nil,
	}
	return &b, nil

}
