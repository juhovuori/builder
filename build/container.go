package build

// Container is the container for builds
type Container interface {
	Init(purge bool) error
	Builds() []string
	Build(ID string) (Build, error)
	New(b Buildable) (Build, error)
	AddStage(buildID string, stage Stage) error
	Output(buildID string, output []byte) error
}

// NewContainer creates a new build container
func NewContainer(t string) (Container, error) {
	var c Container
	switch t {
	case "memory":
		c = memoryContainer{map[string]Build{}}
	case "sqlite":
		c = &sqlContainer{}
	default:
		return nil, ErrContainerType
	}
	err := c.Init(false)
	return c, err
}
