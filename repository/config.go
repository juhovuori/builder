package repository

import uuid "github.com/satori/go.uuid"

// Config is a repository configuration
type Config struct {
	Type Type
	URL  string
}

// ID returns an ID for a repository
func (c Config) ID() string {
	return c.ID()
}

// ID returns an ID for a repository
func ID(t Type, URL string) string {
	return uuid.NewV5(namespace, string(t)+":"+URL).String()
}
