package server

import (
	"log"

	"github.com/mitchellh/cli"
)

// Command is the server command
type Command struct{}

// Help returns long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (cmd *Command) Help() string {
	return "builder server <configuration.hcl>"
}

// Run runs the actual command with the given CLI instance and
// command-line arguments. It returns the exit status when it is
// finished.
func (cmd *Command) Run(args []string) int {
	if len(args) != 1 {
		return cli.RunResultHelp
	}
	filename := args[0]
	cfg, err := DefaultConfig(filename)
	if err != nil {
		log.Println("Unable to read configuration", err.Error)
		return 1
	}
	server, err := New(cfg)
	if err != nil {
		log.Println("Unable to create server", err.Error)
		return 1
	}
	err = server.Run()
	if err != nil {
		log.Println("Unable to run server", err.Error)
		return 1
	}

	return 0
}

// Synopsis returns a one-line, short synopsis of the command.
func (cmd *Command) Synopsis() string {
	return "runs a server"
}

// CommandFactory creates a server command for CLI
func CommandFactory() (cli.Command, error) {
	cmd := Command{}
	return &cmd, nil
}
