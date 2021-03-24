package networking

import (
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

/*
	Keeps our program alive until all threads finish
*/
var stayAlive int32

var HOSTS = make(chan Host)

var PORTS = make(chan Port)

var SCAN_DURATION int
