package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	travelToPast()
	travelToFuture()

	fmt.Println("Back to present")
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}

func travelToPast() {
	time.Sleep(time.Second * 2)
	fmt.Println("Travelling to the past")
}

func travelToFuture() {
	time.Sleep(time.Second * 2)
	fmt.Println("Travelling to the future")
}
