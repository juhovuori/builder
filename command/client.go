package command

import (
	"log"
	"os"

	"github.com/juhovuori/builder/client"
	"github.com/mitchellh/cli"
)

// Client is the client command
type Client struct{}

// Help returns long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (cmd *Client) Help() string {
	return "builder client"
}

// Run runs the actual command with the given CLI instance and
// command-line arguments. It returns the exit status when it is
// finished.
func (cmd *Client) Run(args []string) int {
	url := os.Getenv("BUILDER_URL")
	buildID := os.Getenv("BUILDER_BUILDID")
	client, err := client.New(url, buildID)
	if err != nil {
		log.Printf("Cannot create client: %s\n", err.Error())
		return 1
	}
	if len(args) < 1 {
		log.Println(cmd.Help())
		return 1
	}
	switch args[0] {
	case "shutdown":
		token := os.Getenv("BUILDER_TOKEN")
		if token == "" {
			log.Println("No build token")
			return 1
		}
		err = client.Shutdown(token)
		if err != nil {
			log.Println("Failed to shutdown", err)
			return 1
		}
	}
	return 0
}

// Synopsis returns a one-line, short synopsis of the command.
func (cmd *Client) Synopsis() string {
	return "communicates with builder"
}

// ClientFactory creates a client command for CLI
func ClientFactory() (cli.Command, error) {
	cmd := Client{}
	return &cmd, nil
}
