package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func cli() ([]string, error) {
	printHeader()
	targetPTR := flag.String("t", "", "Target address or network")
	flag.Parse()
	targets, err := parseNetwork(*targetPTR)
	if err != nil {
		return nil, err
	} else {
		return targets, nil
	}
}

func printHeader() {
	fmt.Printf(" _____     _                 _       _____         _           \n")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("|   | |___| |_ _ _ _ ___ ___| |_ ___|   __|___ ___| |_ ___ ___ \n")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("| | | | -_|  _| | | | . |  _| '_|___|  |  | . | . |   | -_|  _|\n")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("|_|___|___|_| |_____|___|_| |_,_|   |_____|___|  _|_|_|___|_|  \n")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("                                              |_|              \n")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("Author: Eric Marcantonio (@EricMarcantonio)\n")
	time.Sleep(time.Millisecond * 200)
}

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
	if ipCIDR.Mask == nil{
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
