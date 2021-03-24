package networking

import (
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

var COMMON_PORTS = []int{
	21,  //FTP
	22,  //SSH
	25,  //SMTP
	53,  //DNS
	80,  //HTTP
	110, //POP
	123, //NTP
	143, //IMAP
	443, //HTTPS
	465, //SMTPS
	631, //CUPS
	993, //IMAPS
	995, //POP3 (secure)
}

func ScanAllPorts(addr *net.IPAddr) {
	for _, port := range COMMON_PORTS {
		atomic.AddInt32(&stayAlive, 1)
		go scanPort(addr, port)
	}
	atomic.AddInt32(&stayAlive, -1)
}

func scanPort(addr *net.IPAddr, port int) {

	thisPort := Port{
		addr: addr,
		num:  port,
	}
	con, err := net.DialTimeout("tcp", net.JoinHostPort(addr.String(), strconv.Itoa(port)), 2*time.Second)
	if err != nil {
		thisPort.isOpen = false
		PORTS <- thisPort
		atomic.AddInt32(&stayAlive, -1)
		//log.Printf("Tried %s but got this: %s", net.JoinHostPort(addr.String(), strconv.Itoa(port)), err.Error())
	} else {
		if con != nil {
			thisPort.isOpen = true
			PORTS <- thisPort
		} else {
			thisPort.isOpen = false
			PORTS <- thisPort
		}
		atomic.AddInt32(&stayAlive, -1)
	}
}
