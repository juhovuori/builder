package build

import "strings"

// Container is the container for builds
type Container interface {
	Init() error
	Close() error
	Purge() error
	Builds(projectID *string) []string
	Build(ID string) (Build, error)
	New(b Buildable) (Build, error)
	AddStage(buildID string, stage Stage) error
	Output(buildID string, output []byte) error
}

// NewContainer creates a new build container
func NewContainer(desc string) (Container, error) {
	var c Container
	parts := strings.SplitN(desc, ":", 2)
	t := parts[0]
	cfg := ""
	if len(parts) == 2 {
		cfg = parts[1]
	}
	switch t {
	case "memory":
		c = memoryContainer{map[string]Build{}}
	case "sqlite":
		if cfg == "" {
			cfg = "/tmp/builder.db"
		}
		c = &sqlContainer{filename: cfg}
	default:
		return nil, ErrContainerType
	}
	err := c.Init()
	return c, err
}
