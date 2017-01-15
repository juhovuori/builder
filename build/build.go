package build

import "github.com/satori/go.uuid"

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
	AddStage(stage Stage) error
	Output([]byte) error
	Stdout() []byte
}

type defaultBuild struct {
	BID           string  `json:"id"`
	BProjectID    string  `json:"project-id"`
	BScript       string  `json:"script"`
	BExecutorType string  `json:"executor-type"`
	BErr          error   `json:"error"`
	Bstages       []Stage `json:"stages"`
	Boutput       []byte
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
	if len(b.Bstages) == 0 {
		return false
	}
	lastStage := b.Bstages[len(b.Bstages)-1]
	return lastStage.Type == SUCCESS || lastStage.Type == FAILURE
}

func (b *defaultBuild) Error() error {
	return b.BErr
}

func (b *defaultBuild) AddStage(stage Stage) error {
	var previousStage *Stage
	if len(b.Bstages) != 0 {
		previousStage = &b.Bstages[len(b.Bstages)-1]
	}
	err := stage.ValidateWithPredecessor(previousStage)
	if err != nil {
		return err
	}
	b.Bstages = append(b.Bstages, stage)
	return nil
}

func (b *defaultBuild) Output(output []byte) error {
	b.Boutput = append(b.Boutput, output...)
	return nil
}

func (b *defaultBuild) Stdout() []byte {
	return b.Boutput
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
	id := uuid.NewV4().String()
	b := defaultBuild{
		BID:           id,
		BProjectID:    project.ID(),
		BScript:       project.Script(),
		BExecutorType: e,
		BErr:          nil,
		Bstages:       []Stage{},
	}
	return &b, nil
}
