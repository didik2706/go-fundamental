package main

import (
	"fmt"
)

func main() {
	total := 0

	for num := 1; num <= 10; num++ {
		total += num
	}

	fmt.Println(total);
}