package networking

import (
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"sync/atomic"
	"time"
)

/*
	Sends host along chan HOSTS
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
		//Change each host before sending it along to the host
		thisHost.rtt = duration
		thisHost.resolved = true
		HOSTS <- thisHost
	}

	// Every host will end up here, just some will fail to resolve. Don't want to fail
	host.OnIdle = func() {
		if !thisHost.resolved {
			HOSTS <- thisHost
		}
		atomic.AddInt32(&stayAlive, -1)
	}

	err = host.Run()
	if err != nil {
		log.Println(err)
		atomic.AddInt32(&stayAlive, -1)
	}

}
