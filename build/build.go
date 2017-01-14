package build

// Buildable can be built interface {
type Buildable interface {
	ID() string
	Script() string
}

// State is the state of a build
type State string

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
	BID           string `json:"id"`
	BProjectID    string `json:"project-id"`
	BScript       string `json:"script"`
	BExecutorType string `json:"executor-type"`
	BCompleted    bool   `json:"completed"`
	BErr          error  `json:"error"`
}

func (b *defaultBuild) ID() string {
	return b.BID
}

func (b *defaultBuild) ExecutorType() string {
	return b.BExecutorType
}

func (b *defaultBuild) ProjectID() string {
	return b.BProjectID
}

func (b *defaultBuild) Script() string {
	return b.BScript
}

func (b *defaultBuild) Completed() bool {
	return b.BCompleted
}

func (b *defaultBuild) Error() error {
	return b.BErr
}

// New creates a new build
func New(project Buildable) (Build, error) {
	if project == nil {
		return nil, ErrNilProject
	}
	e := "nop"
	if project.Script() != "" {
		e = "fork"
	}
	b := defaultBuild{
		BID:           "",
		BProjectID:    project.ID(),
		BScript:       project.Script(),
		BExecutorType: e,
		BCompleted:    false,
		BErr:          nil,
	}
	return &b, nil

}
