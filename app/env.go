package app

import (
	"fmt"
	"log"
	"os"

	"github.com/kardianos/osext"
)

func createEnv(builderURL, buildID string) []string {
	mypath, err := osext.ExecutableFolder()
	if err != nil {
		log.Println("Cannot figure out executable folder. Build path may be incorrect.")
	}
	ospath := os.Getenv("PATH")
	path := fmt.Sprintf("%s:%s", mypath, ospath)
	env := []string{
		fmt.Sprintf("PATH=%s", path),
		fmt.Sprintf("BUILDER_URL=%s", builderURL),
		fmt.Sprintf("BUILDER_BUILD_ID=%s", buildID),
	}
	return append(os.Environ(), env...)
}
