package project

// Manager represents the manager of a single project
type Manager interface {
	URL() string
	MD5() string
	Err() error
	Watched() bool
}
