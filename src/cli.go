package main

import (
	"flag"
	"github.com/gookit/color"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func cli() ([]string, error) {
	targetPTR := flag.String("t", "", "Target address or network")
	flag.Parse()
	if *targetPTR == "" {
		log.Fatalln("No Argument Passed")
	}
	printTimedHeader()
	setTerminalClear()

	targets, err := parseNetwork(*targetPTR)
	if err != nil {
		return nil, err
	} else {
		return targets, nil
	}
}

func printTimedHeader() {
	color.Green.Printf(" _____     _                 _       _____         _           \n")
	time.Sleep(time.Millisecond * 200)
	color.Green.Printf("|   | |___| |_ _ _ _ ___ ___| |_ ___|   __|___ ___| |_ ___ ___ \n")
	time.Sleep(time.Millisecond * 200)
	color.Green.Printf("| | | | -_|  _| | | | . |  _| '_|___|  |  | . | . |   | -_|  _|\n")
	time.Sleep(time.Millisecond * 200)
	color.Green.Printf("|_|___|___|_| |_____|___|_| |_,_|   |_____|___|  _|_|_|___|_|  \n")
	time.Sleep(time.Millisecond * 200)
	color.Green.Printf("                                              |_|              \n")
	time.Sleep(time.Millisecond * 200)
	color.Red.Printf("Author: Eric Marcantonio (@EricMarcantonio)\n")
	time.Sleep(time.Millisecond * 200)
}

func printHeader() {
	color.Green.Printf(" _____     _                 _       _____         _           \n")
	color.Green.Printf("|   | |___| |_ _ _ _ ___ ___| |_ ___|   __|___ ___| |_ ___ ___ \n")
	color.Green.Printf("| | | | -_|  _| | | | . |  _| '_|___|  |  | . | . |   | -_|  _|\n")
	color.Green.Printf("|_|___|___|_| |_____|___|_| |_,_|   |_____|___|  _|_|_|___|_|  \n")
	color.Green.Printf("                                              |_|              \n")
	color.Red.Printf("Author: Eric Marcantonio (@EricMarcantonio)\n")
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

func logFatalf(err string) {

}
