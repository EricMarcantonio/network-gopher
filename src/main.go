package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"
)

var hosts = make(chan Host)
var ports = make(chan Port)

var stayAlive int32

func main() {
	start := time.Now()
	final := make(map[string][]int)
	for i := 0; i < 256; i++ {
		atomic.AddInt32(&stayAlive, 1)
		go TestHost("10.0.0." + strconv.Itoa(i))
	}

	for atomic.LoadInt32(&stayAlive) > 0 {
		select {
		case msg := <-hosts:
			atomic.AddInt32(&stayAlive, -1)
			if msg.resolved {
				atomic.AddInt32(&stayAlive, 1)
				//log.Printf("Found a host up at %s", msg.addr.String())
				go ScanAllPorts(msg.addr)
			}
		case p := <-ports:
			atomic.AddInt32(&stayAlive, -1)
			if p.isOpen {
				//log.Println(net.JoinHostPort(p.addr.String(), strconv.Itoa(p.num)) + "is" + strconv.FormatBool(p.isOpen))
				final[p.addr.String()] = append(final[p.addr.String()], p.num)
			}
		}
	}
	fmt.Println("Address\t\tPorts Up")
	for s, ints := range final {
		fmt.Printf("%s\t%d\n", s, ints)
	}
	fmt.Printf("Scan finished in %d second(s)\n", time.Now().Second()-start.Second())

}
