package server

type Server interface {
	Run() error
}

func New(cfg Config) (Server, error) {
	return nil, nil
}
