package repository

// Repository represents a VCS repository
type Repository interface {
	Type() Type
	URL() string
	Update() error
	Init() error
	Cleanup() error
}

// New creates a new Repository
func New(t Type, URL string) (Repository, error) {
	switch t {
	case git:
		return &gitRepository{URL}, nil
	default:
		return nil, ErrInvalidType
	}
}
