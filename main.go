package main

import (
	"github/EricMarcantonio/network-gopher/cli"
	"github/EricMarcantonio/network-gopher/errors"
	"github/EricMarcantonio/network-gopher/networking"
)

func main() {
	targets, err := cli.Cli()
	errors.CheckErr(err)
	finals, t := networking.ScanNetwork(targets)
	cli.PrettyPrintMap(finals)
	cli.TimeTaken(t)
}
