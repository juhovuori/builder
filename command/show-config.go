package command

import (
	"log"

	"github.com/juhovuori/builder/app"
	"github.com/mitchellh/cli"
)

// ShowConfig is the nop command
type ShowConfig struct{}

// Help returns long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (cmd *ShowConfig) Help() string {
	return "builder show-config <file.hcl>"
}

// Run runs the actual command with the given CLI instance and
// command-line arguments. It returns the exit status when it is
// finished.
func (cmd *ShowConfig) Run(args []string) int {
	if len(args) != 1 {
		log.Println(cmd.Help())
		return 1
	}
	cfg, err := app.NewConfig(args[0])
	if err != nil {
		log.Println("Unable to read configuration", err.Error())
		return 1
	}

	log.Printf("Server address %+v\n", cfg.ServerAddress())
	log.Printf("Projects %+v\n", cfg.Projects())
	log.Printf("Build %+v\n", cfg.Store())
	return 0
}

// Synopsis returns a one-line, short synopsis of the command.
func (cmd *ShowConfig) Synopsis() string {
	return "displays configuration read from a file"
}

// ShowConfigFactory creates a nop command for CLI
func ShowConfigFactory() (cli.Command, error) {
	cmd := ShowConfig{}
	return &cmd, nil
}
