package networking

import (
	"sync/atomic"
	"time"
)

func ScanNetwork(targets []string) (map[string][]int, int) {
	START := time.Now()
	FINAL := make(map[string][]int)
	for _, target := range targets {
		atomic.AddInt32(&stayAlive, 1)
		go TestHost(target)
	}
	for atomic.LoadInt32(&stayAlive) > 0 {
		//fmt.Println(stayAlive)
		select {
		case host := <-HOSTS:
			if host.resolved {
				atomic.AddInt32(&stayAlive, 1)
				go ScanAllPorts(host.addr)
			}
		case p := <-PORTS:
			if p.isOpen {
				FINAL[p.addr.String()] = append(FINAL[p.addr.String()], p.num)
				//cli.ClearTermial()
				//cli.PrintHeader()
				//PrettyPrintFinal()
			}
		}
	}
	SCAN_DURATION = time.Now().Second() - START.Second()
	return FINAL, SCAN_DURATION
	//We are done scanning and we know what ports are up. What should we do?
}
