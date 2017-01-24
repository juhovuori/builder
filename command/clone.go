package command

import (
	"log"
	"os"

	"github.com/juhovuori/builder/repository"
	"github.com/mitchellh/cli"
)

// Clone is the clone command
type Clone struct{}

// Help returns long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (cmd *Clone) Help() string {
	return "builder clone"
}

// Run runs the actual command with the given CLI instance and
// command-line arguments. It returns the exit status when it is
// finished.
func (cmd *Clone) Run(args []string) int {
	var path string
	t := os.Getenv("BUILDER_REPOSITORY_TYPE")
	url := os.Getenv("BUILDER_REPOSITORY")
	r, err := repository.New(repository.Type(t), url)
	if err != nil {
		log.Printf("Cannot instantiate repository: %s\n", err.Error())
		log.Printf("Check your BUILDER_REPOSITORY_TYPE and BUILDER_REPOSITORY variables\n")
		return 1
	}
	if len(args) > 0 {
		path = args[0]
	} else {
		path, err = os.Getwd()
		if err != nil {
			log.Println("Failed to figure out working directory", err)
			return 1
		}
	}
	err = r.Clone(path)
	if err != nil {
		log.Println("Failed to clone", err)
		return 1
	}
	return 0
}

// Synopsis returns a one-line, short synopsis of the command.
func (cmd *Clone) Synopsis() string {
	return "clone a project"
}

// CloneFactory creates a clone command for CLI
func CloneFactory() (cli.Command, error) {
	cmd := Clone{}
	return &cmd, nil
}
