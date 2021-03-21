package main

import (
	"flag"
	"fmt"
	"github.com/google/goterm/term"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func cli() ([]string, error) {
	printTimedHeader()
	setTerminalClear()
	targetPTR := flag.String("t", "", "Target address or network")
	flag.Parse()
	targets, err := parseNetwork(*targetPTR)
	if err != nil {
		return nil, err
	} else {
		return targets, nil
	}
}

func printTimedHeader() {
	fmt.Printf(term.Greenf(" _____     _                 _       _____         _           \n"))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf(term.Greenf("|   | |___| |_ _ _ _ ___ ___| |_ ___|   __|___ ___| |_ ___ ___ \n"))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf(term.Greenf("| | | | -_|  _| | | | . |  _| '_|___|  |  | . | . |   | -_|  _|\n"))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf(term.Greenf("|_|___|___|_| |_____|___|_| |_,_|   |_____|___|  _|_|_|___|_|  \n"))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf(term.Greenf("                                              |_|              \n"))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf(term.Redf("Author: Eric Marcantonio (@EricMarcantonio)\n"))
	time.Sleep(time.Millisecond * 200)
}

func printHeader() {
	fmt.Printf(term.Greenf(" _____     _                 _       _____         _           \n"))
	fmt.Printf(term.Greenf("|   | |___| |_ _ _ _ ___ ___| |_ ___|   __|___ ___| |_ ___ ___ \n"))
	fmt.Printf(term.Greenf("| | | | -_|  _| | | | . |  _| '_|___|  |  | . | . |   | -_|  _|\n"))
	fmt.Printf(term.Greenf("|_|___|___|_| |_____|___|_| |_,_|   |_____|___|  _|_|_|___|_|  \n"))
	fmt.Printf(term.Greenf("                                              |_|              \n"))
	fmt.Printf(term.Redf("Author: Eric Marcantonio (@EricMarcantonio)\n"))
}

func ClearTermial() {
	clearFunc, ok := clear[runtime.GOOS]
	if ok {
		clearFunc()
	} else {
		log.Panicln("I don't know what platform you built this on...but I cannot see your screen")
	}
}

func setTerminalClear() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		checkErr(err)
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		checkErr(err)
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		checkErr(err)
	}
}
