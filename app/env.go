package app

import (
	"fmt"
	"log"
	"os"

	"github.com/juhovuori/builder/repository"
	"github.com/kardianos/osext"
)

func createEnv(r repository.Repository, builderURL, buildID string) []string {
	mypath, err := osext.ExecutableFolder()
	if err != nil {
		log.Println("Cannot figure out executable folder. Build path may be incorrect.")
	}
	ospath := os.Getenv("PATH")
	path := fmt.Sprintf("%s:%s", mypath, ospath)
	env := []string{
		fmt.Sprintf("BUILDER_BUILD_ID=%s", buildID),
		fmt.Sprintf("BUILDER_REPOSITORY=%s", r.URL()),
		fmt.Sprintf("BUILDER_REPOSITORY_TYPE=%s", r.Type()),
		fmt.Sprintf("BUILDER_URL=%s", builderURL),
		fmt.Sprintf("PATH=%s", path),
	}
	return append(os.Environ(), env...)
}
