package main

import (
	"fmt"
	"path"
	"syscall"

	"github.com/coreos/fleet/third_party/github.com/codegangsta/cli"
)

func newCatUnitCommand() cli.Command {
	return cli.Command{
		Name:	"cat",
		Usage:	"Print the contents of a unit that has been loaded in the cluster",
		Action:	printUnitAction,
	}
}

func printUnitAction(c *cli.Context) {
	r := getRegistry(c)

	if len(c.Args()) != 1 {
		fmt.Println("One unit file must be provided.")
		syscall.Exit(1)
	}

	name := path.Base(c.Args()[0])
	payload := r.GetPayload(name)

	if payload == nil {
		fmt.Println("Job not found.")
		syscall.Exit(1)
	}

	fmt.Print(payload.Unit.String())
}