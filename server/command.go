package server

import (
	"flag"
	"log"

	"github.com/juhovuori/builder/app"
	"github.com/mitchellh/cli"
)

// Command is the server command
type Command struct{}

// Help returns long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (cmd *Command) Help() string {
	return "builder server [-f <configuration.hcl>]"
}

// Run runs the actual command with the given CLI instance and
// command-line arguments. It returns the exit status when it is
// finished.
func (cmd *Command) Run(args []string) int {
	var configFile string
	cmdFlags := flag.NewFlagSet("server", flag.ContinueOnError)
	cmdFlags.Usage = func() { log.Println(cmd.Help()) }
	cmdFlags.StringVar((&configFile), "config-file", "./builder.hcl", "HCL file to read config from")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	app, err := app.NewFromFilename(configFile)
	if err != nil {
		log.Println("Unable to start application", err.Error())
		return 1
	}
	server, err := New(app)
	if err != nil {
		log.Println("Unable to create server", err.Error())
		return 1
	}
	err = server.Run()
	if err != nil {
		log.Println("Unable to run server", err.Error())
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
