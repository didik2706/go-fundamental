package main

import (
	"fmt"
)

func main() {
	var members[3] string

	members[0] = "Didik"
	members[1] = "Hafidz"
	members[2] = "Nurul"

	angka := [5]int{1,2,3,4,5}
	angkaLagi := [...]int {
		1,
		2,
		3,
		4,
		5,
		5,
		5,
	}

	// menampilkan isi array menggunakan for loop
	for i := 0; i < len(members); i++ {
		fmt.Println("Membernya adalah " + members[i])
	}

	// menggunakan perulangan range
	for _, value := range members {
		// fmt.Printf("Index ke %d harinya %s \n", index, value);
		fmt.Println(value);
	}

	// array multidimensi
	animals := [2][3]string {
		{ "Macan", "Singa", "Harimau" },
		{ "Elang", "Gagak", "Rajawali" },
	}

	// loop array multidimensi
	for _, animal := range animals {
		for _, value := range animal {
			fmt.Println(value)	
		}
		fmt.Println("----------")
	}

	fmt.Println(angka)
	fmt.Println(angkaLagi)
	fmt.Println(members)
}