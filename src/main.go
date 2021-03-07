package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"strconv"
	"time"
)

type Host struct {
	addr     *net.IPAddr
	rtt      time.Duration
	resolved bool
}

var responses = make(chan Host, 20)

func main() {
	start := time.Now()
	for i := 0; i < 256; i++ {
		go pingHost("10.0.0." + strconv.Itoa(i))
	}
	i := 0
	for {
		msg := <-responses
		log.Println(strconv.Itoa(i) + ": " + msg.addr.String() + " in " + msg.rtt.String())
		i++
		if i == 256 {
			break
		}
	}
	fmt.Println("Scanned " + strconv.Itoa(i) + " ipaddresses in " + strconv.Itoa(time.Now().Second() - start.Second()) + "s.")

}

func pingHost(addr string) {

	ping := fastping.NewPinger()
	ip, err := net.ResolveIPAddr("ip4:icmp", addr)

	if err != nil {
		log.Fatalln(err.Error())
	}
	HOST := Host{
		addr: ip,
		rtt: 0,
		resolved: false,
	}
	ping.AddIPAddr(ip)
	ping.OnRecv = func(addr *net.IPAddr, duration time.Duration) {
		HOST.resolved = true
		HOST.rtt = duration
		responses <- HOST
	}
	ping.OnIdle = func() {
		if !HOST.resolved {
			responses <- HOST
		}
	}
	err = ping.Run()
	if err != nil {
		log.Println(err)
	}
}
