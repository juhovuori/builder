package repository

type nopRepository struct {
	id  string
	url string
}

func (r *nopRepository) Type() Type {
	return nop
}

func (r *nopRepository) URL() string {
	return r.url
}

func (r *nopRepository) ID() string {
	return r.id
}

func (r *nopRepository) File(filename string) ([]byte, error) {
	return []byte{}, nil
}

func (r *nopRepository) Init() error {
	return nil
}

func (r *nopRepository) Cleanup() error {
	return nil
}

func (r *nopRepository) Update() error {
	return nil
}
