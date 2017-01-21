package command

import (
	"log"
	"os"

	"github.com/juhovuori/builder/client"
	"github.com/mitchellh/cli"
)

// Build is the build command
type Build struct{}

// Help returns long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (cmd *Build) Help() string {
	return "builder build"
}

// Run runs the actual command with the given CLI instance and
// command-line arguments. It returns the exit status when it is
// finished.
func (cmd *Build) Run(args []string) int {
	url := os.Getenv("BUILDER_URL")
	client, err := client.NewWithoutBuildID(url)
	if err != nil {
		log.Printf("Cannot create client: %s\n", err.Error())
		return 1
	}
	if len(args) < 1 {
		log.Println(cmd.Help())
		return 1
	}
	projectID := args[0]
	build, err := client.Build(projectID)
	if err != nil {
		log.Println("Failed to build", err)
		return 1
	}
	log.Println(build)
	return 0
}

// Synopsis returns a one-line, short synopsis of the command.
func (cmd *Build) Synopsis() string {
	return "build a project"
}

// BuildFactory creates a build command for CLI
func BuildFactory() (cli.Command, error) {
	cmd := Build{}
	return &cmd, nil
}
