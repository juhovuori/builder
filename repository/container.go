package repository

import "github.com/labstack/gommon/log"

// Container is the repositories manager
type Container interface {
	Repositories() []string
	Repository(t Type, URL string) (Repository, error)
	Add(t Type, URL string) error
	Remove(t Type, URL string) error
}

type defaultContainer struct {
	repositories map[string]Repository
}

func (c *defaultContainer) Repositories() []string {
	repositories := []string{}
	for ID := range c.repositories {
		repositories = append(repositories, ID)
	}
	return repositories
}

func (c *defaultContainer) Repository(Type Type, URL string) (Repository, error) {
	id := ID(Type, URL)
	p, ok := c.repositories[id]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (c *defaultContainer) Add(t Type, URL string) error {
	repository, err := New(t, URL)
	if err != nil {
		return err
	}
	id := repository.ID()
	if _, ok := c.repositories[id]; ok {
		return ErrDuplicate
	}
	c.repositories[id] = repository
	return nil
}

func (c *defaultContainer) Remove(t Type, URL string) error {
	id := ID(t, URL)
	r, ok := c.repositories[id]
	if !ok {
		return ErrNotFound
	}
	err := r.Cleanup()
	if err != nil {
		log.Errorf("Error cleaning up %s %s: %v\n", t, URL, err)
	}
	delete(c.repositories, id)
	return nil
}

// NewContainer creates a new repository manager
func NewContainer() Container {
	c := &defaultContainer{map[string]Repository{}}
	return c
}
