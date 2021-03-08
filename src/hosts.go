package main

import (
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"time"
)

/*
Represents a host on the network.
*/

type Host struct {
	/*
		The IP of the host
	*/
	addr *net.IPAddr
	/*
		The time it takes to get from host to target and back.
	*/
	rtt time.Duration
	/*
		Whether or not the address is up.
		Note for the security folks: I am lying to you...
	*/
	resolved bool
}

/*
	Sends host along chan `hosts`
*/
func TestHost(addr string) {
	thisHost := Host{
		resolved: false,
	}
	host := fastping.NewPinger()
	ip, err := net.ResolveIPAddr("ip4:icmp", addr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	host.AddIPAddr(ip)
	thisHost.addr = ip
	host.OnRecv = func(addr *net.IPAddr, duration time.Duration) {
		thisHost.rtt = duration
		thisHost.resolved = true
		hosts <- thisHost
	}

	host.OnIdle = func() {
		if !thisHost.resolved {
			hosts <- thisHost
		}
	}

	err = host.Run()
	if err != nil {
		log.Println(err)
	}
}
