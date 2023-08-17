package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func main() {

	start := time.Now()

	go TravelToUniverse("Echoverse")
	go TravelToUniverse("Nexusverse")
	go TravelToUniverse("Mirrordimension")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Back to Origin")
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}

func TravelToUniverse(universe string) {
	time.Sleep(time.Second * 2)
	fmt.Printf("Travelling to %s Universe\n", universe)
	ch <- fmt.Sprintf("Back from %s Universe\n", universe)
}
