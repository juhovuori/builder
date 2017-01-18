package command

import (
	"flag"
	"log"
	"os"

	"github.com/juhovuori/builder/app"
	"github.com/juhovuori/builder/server"
	"github.com/mitchellh/cli"
)

// Server is the server command
type Server struct{}

// Help returns long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (cmd *Server) Help() string {
	return "builder server [-f <configuration.hcl>]"
}

// Run runs the actual command with the given CLI instance and
// command-line arguments. It returns the exit status when it is
// finished.
func (cmd *Server) Run(args []string) int {
	var configFile string
	var systemToken *string
	cmdFlags := flag.NewFlagSet("server", flag.ContinueOnError)
	cmdFlags.Usage = func() { log.Println(cmd.Help()) }
	cmdFlags.StringVar((&configFile), "f", "./builder.hcl", "HCL file to read config from")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if configFile == "" {
		configFile = os.Getenv("BUILDER_CONFIG")
	}
	if configFile == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Println("Unable to get working directory", err.Error())
			return 1
		}
		configFile = wd + "/builder.hcl"
	}

	rawSystemToken := os.Getenv("BUILDER_TOKEN")
	if rawSystemToken != "" {
		systemToken = &rawSystemToken
	}

	app, err := app.NewFromURL(configFile)
	if err != nil {
		log.Println("Unable to start application", err.Error())
		return 1
	}
	s, err := server.NewWithSystemToken(app, systemToken)
	if err != nil {
		log.Println("Unable to create server", err.Error())
		return 1
	}
	err = s.Run()
	if err != nil {
		log.Println("Unable to run server", err.Error())
		return 1
	}

	return 0
}

// Synopsis returns a one-line, short synopsis of the command.
func (cmd *Server) Synopsis() string {
	return "runs a server"
}

// ServerFactory creates a server command for CLI
func ServerFactory() (cli.Command, error) {
	cmd := Server{}
	return &cmd, nil
}
