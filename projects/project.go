package projects

import (
	"crypto/md5"
	"fmt"
)

// Project represents a single project
type Project struct {
	URL string
	MD5 string
	Err error
	Cfg struct{}
}

func newProject(URL string) Project {
	MD5 := md5.Sum([]byte(URL))
	p := Project{
		URL: URL,
		MD5: fmt.Sprintf("%x", MD5),
		Err: ErrFetchProject,
	}

	return p
}
