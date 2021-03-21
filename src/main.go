package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

/*
Represents a host on the network.
*/

var hosts = make(chan Host)
var ports = make(chan Port)
var stayAlive int32
var final map[string][]int

func main() {
	targets, err := cli()
	if err != nil {
		log.Fatalln(err.Error())
	}
	start := time.Now()
	final = make(map[string][]int)
	for _, target := range targets {
		atomic.AddInt32(&stayAlive, 1)
		go TestHost(target)
	}


	for atomic.LoadInt32(&stayAlive) > 0 {
		//fmt.Println(stayAlive)
		select {
		case host := <-hosts:
			if host.resolved {
				atomic.AddInt32(&stayAlive, 1)
				go ScanAllPorts(host.addr)
			}
		case p := <-ports:
			if p.isOpen {
				final[p.addr.String()] = append(final[p.addr.String()], p.num)
			}
		}
	}
	fmt.Printf("Scan finished in %d second(s)\n", time.Now().Second()-start.Second())
	PrettyPrintFinal()
}

func PrettyPrintFinal() {
	fmt.Printf("IP\t\tPorts\n")
	for ip, ports := range final {
		fmt.Printf("%s\t%d\n", ip, ports)
	}
	fmt.Printf("Back to hiding!! \n")
}
