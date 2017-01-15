package exec

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"sync"
	"time"
)

var once sync.Once

func tmpFilename() string {
	once.Do(func() { rand.Seed(time.Now().UnixNano()) })
	n := rand.Int63()
	dir := path.Join(os.TempDir(), fmt.Sprintf("builder-%019d", n))
	return dir
}

func tmpFilenameByID(id string) string {
	dir := path.Join(os.TempDir(), fmt.Sprintf("builder-%s", id))
	return dir
}
