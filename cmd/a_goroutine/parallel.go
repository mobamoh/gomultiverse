package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	go travel("Past")
	go travel("Future")

	time.Sleep(time.Second * 3)

	fmt.Println("Back to present")
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}

func travel(direction string) {
	time.Sleep(time.Second * 2)
	fmt.Printf("Travelling to the %s...\n", direction)
}
