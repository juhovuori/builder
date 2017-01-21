package repository

import (
	"fmt"
	"os"
	"path"
)

func tmpFilenameByID(id string) string {
	dir := path.Join(os.TempDir(), fmt.Sprintf("builder-repo-%s", id))
	return dir
}
