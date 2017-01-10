package main

import (
	"log"
	"os"

	"github.com/juhovuori/builder/command"
	"github.com/juhovuori/builder/server"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("builder", "0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"nop":    command.NopFactory,
		"server": server.CommandFactory,
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
