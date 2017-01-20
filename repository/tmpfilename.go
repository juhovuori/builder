package repository

import (
	"fmt"
	"os"
	"path"
)

func tmpFilenameByID(id string) string {
	dir := path.Join(os.TempDir(), fmt.Sprintf("builder-%s", id))
	return dir
}
