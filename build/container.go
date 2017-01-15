package build

// Container is the container for builds
type Container interface {
	Builds() []string
	Build(ID string) (Build, error)
	New(b Buildable) (Build, error)
	AddStage(buildID string, stage Stage) error
	Output(buildID string, output []byte) error
}

type memoryContainer struct {
	builds map[string]Build
}

func (c memoryContainer) Builds() []string {
	builds := []string{}
	for ID := range c.builds {
		builds = append(builds, ID)
	}
	return builds
}

func (c memoryContainer) Build(ID string) (Build, error) {
	build, ok := c.builds[ID]
	if !ok {
		return nil, ErrNotFound
	}
	return build, nil
}

func (c memoryContainer) New(b Buildable) (Build, error) {
	build, err := New(b)
	if err != nil {
		return nil, err
	}
	c.builds[build.ID()] = build
	return build, nil
}

func (c memoryContainer) AddStage(buildID string, stage Stage) error {
	b, err := c.Build(buildID)
	if err != nil {
		return err
	}
	return b.AddStage(stage)
}

func (c memoryContainer) Output(buildID string, output []byte) error {
	b, err := c.Build(buildID)
	if err != nil {
		return err
	}
	return b.Output(output)
}

// NewContainer creates a new build container
func NewContainer(t string) (Container, error) {
	switch t {
	case "memory":
		return memoryContainer{map[string]Build{}}, nil
	default:
		return nil, ErrContainerType
	}
}
