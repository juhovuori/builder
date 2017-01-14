package build

// Container is the container for builds
type Container interface {
	Builds() []string
	Build(ID string) (Build, error)
	New(b Buildable) (Build, error)
	AddStage(ID string, state State) error
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

func (c memoryContainer) AddStage(ID string, state State) error {
	return errNotImplemented
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
