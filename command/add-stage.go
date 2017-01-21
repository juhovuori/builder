package command

import (
	"log"
	"os"

	"github.com/juhovuori/builder/client"
	"github.com/mitchellh/cli"
)

// AddStage is the add-stage command
type AddStage struct{}

// Help returns long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (cmd *AddStage) Help() string {
	return "builder add-stage"
}

// Run runs the actual command with the given CLI instance and
// command-line arguments. It returns the exit status when it is
// finished.
func (cmd *AddStage) Run(args []string) int {
	url := os.Getenv("BUILDER_URL")
	buildID := os.Getenv("BUILDER_BUILD_ID")
	client, err := client.New(url, buildID)
	if err != nil {
		log.Printf("Cannot create client: %s\n", err.Error())
		return 1
	}
	if len(args) < 1 {
		log.Println(cmd.Help())
		return 1
	}
	stage := args[0]
	if err = client.AddStage(stage, os.Stdin); err != nil {
		log.Println("Failed to add stage", err)
		return 1
	}
	return 0
}

// Synopsis returns a one-line, short synopsis of the command.
func (cmd *AddStage) Synopsis() string {
	return "add a build stage"
}

// AddStageFactory creates a add-stage command for CLI
func AddStageFactory() (cli.Command, error) {
	cmd := AddStage{}
	return &cmd, nil
}
