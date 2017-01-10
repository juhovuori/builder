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
		"nop":    command.NopFactory,
		"server": command.ServerFactory,
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
