package main

import (
	"fmt"
)

func main() {
	members := [5]string { "Didik", "Hafidz", "Nurul", "Sam", "Dzun" }

	// membuat slice menggunakan keyword make
	makeSlice := make([]string, 5, 10)
	copy(makeSlice, members[:])
	makeSlice = append(makeSlice, "Karim", "Maman")

	// mmebuat slice
	newSlice := members[:2]
	newSlice[1] = "Karim"


	// menthod untuk slice
	fmt.Println("Panjang slice", len(makeSlice))
	fmt.Println("kapasitas slice", cap(makeSlice))

	fmt.Println(members)
	fmt.Println(newSlice)
	fmt.Println(makeSlice)
}