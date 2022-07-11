package main

import (
	"fmt"
	"strconv"
)


func main() {
	// ini adalah komentar

	gaji := "10000000"
	gajiconv, _ := strconv.Atoi(gaji);
	bonus := 20000 + gajiconv
	// convert from int to string menggunakan Itoa sebalikannya menggunak Atoi

	fmt.Println("total " + strconv.Itoa(bonus))
}