package main

import (
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

/*
	Represents a port on a computer
*/
type Port struct {
	/*
		The address for this port
	*/
	addr *net.IPAddr
	/*
		The port number
	*/
	num int
	/*
		Open status of the port
	*/
	isOpen bool
}

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
	con, err := net.DialTimeout("tcp", net.JoinHostPort(addr.String(), strconv.Itoa(port)), 2*time.Second)
	if err != nil {
		ports <- Port{
			addr:   addr,
			num:    port,
			isOpen: false, //TODO bug
		}
		//log.Printf("Tried %s but got this: %s", net.JoinHostPort(addr.String(), strconv.Itoa(port)), err.Error())
	} else {
		if con != nil {
			ports <- Port{
				addr:   addr,
				num:    port,
				isOpen: true, //TODO bug
			}
		} else {
			ports <- Port{
				addr:   addr,
				num:    port,
				isOpen: false,
			}
		}
	}
}
