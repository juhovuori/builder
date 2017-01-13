package project

// Config is the configuration for projects container.
type Config interface {
	Projects() []string
}
