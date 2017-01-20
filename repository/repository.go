package repository

import uuid "github.com/satori/go.uuid"

// Repository represents a VCS repository
type Repository interface {
	ID() string
	Type() Type
	URL() string
	Update() error
	Init() error
	Cleanup() error
}

// New creates a new Repository
func New(t Type, URL string) (Repository, error) {

	ID := ID(t, URL)

	switch t {
	case git:
		dir := tmpFilenameByID(ID)
		return &gitRepository{ID, URL, dir}, nil
	case nop:
		return &nopRepository{ID, URL}, nil
	default:
		return nil, ErrInvalidType
	}
}

// namespace is a global namespace for repository ID generation.
var namespace, _ = uuid.FromString("7526818e-fea1-41c4-9f43-870d0c3da3ee")
