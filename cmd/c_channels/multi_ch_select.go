package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	chAndromeda := make(chan string)
	chMilkyWay := make(chan string)
	chSombrero := make(chan string)

	start := time.Now()

	go SignalFromGalaxy("Andromeda", time.Second*1, chAndromeda)
	go SignalFromGalaxy("Milky Way", time.Second*3, chMilkyWay)
	go SignalFromGalaxy("Sombrero", time.Second*9, chSombrero)

loop:
	for {
		select {
		case galaxy := <-chAndromeda:
			fmt.Println(galaxy)
		case galaxy := <-chMilkyWay:
			fmt.Println(galaxy)
		case galaxy := <-chSombrero:
			fmt.Println(galaxy)
			break loop
		}
	}

	fmt.Printf("Elapsed time: %s\n", time.Since(start))
	os.Exit(0)
}

func SignalFromGalaxy(galaxy string, duration time.Duration, ch chan<- string) {
	for {
		time.Sleep(duration)
		ch <- "Signal from " + galaxy
	}
}
