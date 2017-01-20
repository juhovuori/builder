package repository

// Type identifies a repository type
type Type string

const (
	git Type = "git"
)

// Validate validates a repository type
func (t Type) Validate() error {
	switch t {
	case git:
		return nil
	default:
		return ErrInvalidType
	}
}
