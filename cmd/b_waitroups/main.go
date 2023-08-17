package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	start := time.Now()
	go travelToGalaxy("Cigar")
	go travelToGalaxy("Black Eye")

	wg.Wait()
	fmt.Println("Back to Origin")
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}

func travelToGalaxy(galaxy string) {
	time.Sleep(time.Second * 2)
	fmt.Printf("Travelling to %s Galaxy ...\n", galaxy)
	wg.Done()
}
