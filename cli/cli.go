package cli

import (
	"flag"
	"github/EricMarcantonio/network-gopher/networking"
	"log"
)

func Cli() ([]string, error) {
	setTerminalClear()
	targetPTR := flag.String("t", "", "Target address or network")
	flag.Parse()
	if *targetPTR == "" {
		log.Fatalln("No Argument Passed")
	}
	ClearTermial()
	PrintTimedHeader()

	targets, err := networking.ParseNetwork(*targetPTR)
	if err != nil {
		return nil, err
	} else {
		return targets, nil
	}
}
