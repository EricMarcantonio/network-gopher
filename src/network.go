package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

func parseNetwork(target string) ([]string, error) {
	var hostSlice []string
	_, ipCIDR, err := net.ParseCIDR(target)
	if err != nil {
		ipCIDR = &net.IPNet{IP: net.ParseIP(target)}
		if ipCIDR.IP == nil {
			log.Fatal("incorrect ip passed")
		}
	}
	var subnets []string
	if ipCIDR.Mask == nil {
		subnets = strings.Split(ipCIDR.IP.String(), ".")
	} else {
		subnets = strings.Split(ipCIDR.String(), ".")
	}

	isItCIDR := strings.Index(subnets[3], "/")
	if isItCIDR > 0 {
		//They gave us a slash - we have a network to look at
		lastSub := subnets[3][isItCIDR+1:]

		mask, _ := strconv.Atoi(lastSub)
		if mask == 24 {
			for i := 0; i < 256; i++ {
				hostSlice = append(hostSlice, subnets[0]+"."+subnets[1]+"."+subnets[2]+"."+strconv.Itoa(i))
			}
		}
		return hostSlice, nil
	} else {
		hostSlice = append(hostSlice, target)
		return hostSlice, nil
	}
}
