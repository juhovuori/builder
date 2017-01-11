package project

// Manager represents the manager of a single project
type Manager interface {
	URL() string
	ID() string
	Err() error
}
