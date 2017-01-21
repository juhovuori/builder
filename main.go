package main

import (
	"log"
	"os"

	"github.com/juhovuori/builder/command"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("builder", "0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"shutdown":    command.ShutdownFactory,
		"add-stage":   command.AddStageFactory,
		"build":       command.BuildFactory,
		"nop":         command.NopFactory,
		"server":      command.ServerFactory,
		"show-config": command.ShowConfigFactory,
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
