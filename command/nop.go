package command

import "github.com/mitchellh/cli"

// Nop is the nop command
type Nop struct{}

// Help returns long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (cmd *Nop) Help() string {
	return "builder nop"
}

// Run runs the actual command with the given CLI instance and
// command-line arguments. It returns the exit status when it is
// finished.
func (cmd *Nop) Run(args []string) int {
	return 0
}

// Synopsis returns a one-line, short synopsis of the command.
func (cmd *Nop) Synopsis() string {
	return "does nothing"
}

// NopFactory creates a nop command for CLI
func NopFactory() (cli.Command, error) {
	cmd := Nop{}
	return &cmd, nil
}
