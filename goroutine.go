package main

import (
	"fmt"
	"time"
)

func printSalam(text string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Microsecond)
		fmt.Println(text)
	}
}

func main() {
	start := time.Now()

	go printSalam("Assalamualaikum")
	go printSalam("Waalaikumsalam")

	fmt.Println(time.Since(start))
}