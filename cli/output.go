package cli

import (
	"github.com/gookit/color"
)

func PrettyPrintMap(aMap map[string][]int) {
	color.Yellow.Printf("IP\t\tPorts\n")
	for ip, ports := range aMap {
		color.White.Printf("%s\t", ip)
		color.Cyan.Printf("%d\n", ports)
	}
}

func TimeTaken(n int) {
	color.Red.Printf("Scan finished in %d second(s)\n", n)
}
