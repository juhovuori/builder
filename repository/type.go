package repository

// Type identifies a repository type
type Type string

const (
	git Type = "git"
	nop Type = "nop"
)

// Validate validates a repository type
func (t Type) Validate() error {
	switch t {
	case git:
		return nil
	case nop:
		return nil
	default:
		return ErrInvalidType
	}
}
