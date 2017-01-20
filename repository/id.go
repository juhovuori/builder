package repository

import uuid "github.com/satori/go.uuid"

// ID returns an ID for a repository
func ID(t Type, URL string) string {
	return uuid.NewV5(namespace, string(t)+":"+URL).String()
}
