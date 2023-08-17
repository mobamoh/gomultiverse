package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	multiverse := []string{"Echoverse", "Parallax Realm", "Quanta Parallel",
		"Nexusverse", "Doppelgängiverse", "Anomalous Realm",
		"Xverse", "Yverse", "Zrealm",
	}

	var wg sync.WaitGroup
	for i := 0; i < 9; i++ { // This allows us to create MANY goroutines
		wg.Add(1)
		go func(j int) {

			parallelChan := make(chan string)
			originChan := make(chan string)
			travelToUniverse := func() {
				time.Sleep(time.Second * 3) // Deliberate time delay
				parallelChan <- "Travelling to " + multiverse[j]
			}
			signalToOrigin := func() {
				fmt.Println(j, "", <-parallelChan)
				// blocks until it can read something from parallelChan - printed
				originChan <- "Travel Done. Back To Origin from " + multiverse[j] // Done Signal.
			}

			go travelToUniverse()
			go signalToOrigin()
			fmt.Println(j, "", <-originChan)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
