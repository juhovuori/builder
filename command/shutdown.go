package command

import (
	"log"
	"os"

	"github.com/juhovuori/builder/client"
	"github.com/mitchellh/cli"
)

// Shutdown is the shutdown command
type Shutdown struct{}

// Help returns long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (cmd *Shutdown) Help() string {
	return "builder shutdown"
}

// Run runs the actual command with the given CLI instance and
// command-line arguments. It returns the exit status when it is
// finished.
func (cmd *Shutdown) Run(args []string) int {
	url := os.Getenv("BUILDER_URL")
	token := os.Getenv("BUILDER_TOKEN")
	if token == "" {
		log.Println("No build token")
		return 1
	}
	client, err := client.NewWithoutBuildID(url)
	if err != nil {
		log.Printf("Cannot create client: %s\n", err.Error())
		return 1
	}
	if err = client.Shutdown(token); err != nil {
		log.Println("Failed to shutdown", err)
		return 1
	}
	return 0
}

// Synopsis returns a one-line, short synopsis of the command.
func (cmd *Shutdown) Synopsis() string {
	return "shutdown builder"
}

// ShutdownFactory creates a shutdown command for CLI
func ShutdownFactory() (cli.Command, error) {
	cmd := Shutdown{}
	return &cmd, nil
}
