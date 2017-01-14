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
	return nil
}

func (c memoryContainer) Build(ID string) (Build, error) {
	return nil, errNotImplemented
}

func (c memoryContainer) New(b Buildable) (Build, error) {
	return New(b)
}

func (c memoryContainer) AddStage(ID string, state State) error {
	return errNotImplemented
}

// NewContainer creates a new build container
func NewContainer(t string) (Container, error) {
	switch t {
	case "memory":
		return memoryContainer{}, nil
	default:
		return nil, ErrInvalidContainerType
	}
}
