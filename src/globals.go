package main

import "log"

/////////////////////////GLOBAL VARIABLES//////////////////////////////////
/*
Channel to accept hosts after we have worked with them
*/
var hosts = make(chan Host)

/*
Channel to accept each port after we have worked with them
*/
var ports = make(chan Port)

/*
ONLY access with atomic. Used to keep everything alive until its done
*/
var stayAlive int32

/*
The final results
*/
var final map[string][]int

/*
Clear the terminal per platform
*/
var clear map[string]func()

/////////////////////////END//////////////////////////////////
/////////////////////////GLOBAL FUNCTIONS//////////////////////////////////

func checkErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

/////////////////////////END//////////////////////////////////
