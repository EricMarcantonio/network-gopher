package main

import (
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"strconv"
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

type Port struct {
	addr   *net.IPAddr
	num    int
	isOpen bool
}

var hosts = make(chan Host)
var ports = make(chan Port)

func main() {
	final := make(map[string][]int)
	for i := 0; i < 256; i++ {
		go pingHost("10.0.0." + strconv.Itoa(i))
	}
	for {
		select {
		case msg := <-hosts:
			if msg.resolved {
				scanAllPorts(msg.addr)
			}
		case p := <-ports:
			if p.isOpen {
				log.Println(net.JoinHostPort(p.addr.String(), strconv.Itoa(p.num)) + "is" + strconv.FormatBool(p.isOpen))
				final[p.addr.String()] = append(final[p.addr.String()], p.num)
			}
		}
	}
}

func pingHost(addr string) {

	ping := fastping.NewPinger()
	ip, err := net.ResolveIPAddr("ip4:icmp", addr)

	if err != nil {
		log.Fatalln(err.Error())
	}
	HOST := Host{
		addr:     ip,
		rtt:      0,
		resolved: false,
	}
	ping.AddIPAddr(ip)
	ping.OnRecv = func(addr *net.IPAddr, duration time.Duration) {
		HOST.resolved = true
		HOST.rtt = duration
		hosts <- HOST
	}
	ping.OnIdle = func() {
		if !HOST.resolved {
			hosts <- HOST
		}
	}
	err = ping.Run()
	if err != nil {
		log.Println(err)
	}
}

func scanAllPorts(addr *net.IPAddr) {
	common := []int{21, 22, 25, 53, 80, 110, 123, 143, 43, 465, 631, 993, 995}
	for _, i := range common {
		i := i
		go func() {
			_, err := net.Dial("tcp", net.JoinHostPort(addr.String(), strconv.Itoa(i)))
			if err != nil {
				ports <- Port{
					addr:   addr,
					num:    i,
					isOpen: false,
				}
			} else {
				ports <- Port{
					addr:   addr,
					num:    i,
					isOpen: true,
				}
			}
		}()
	}
}
