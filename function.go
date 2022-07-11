package main

import (
	"fmt"
	"strconv"
)

func viewNumber() int {
	return 10;
}

func multiply(num1 int, num2 int) int {
	return num1 * num2
}

func getBiography(umur int, nama string, status string) (string, string) {
	return nama + " adalah seorang " + status,
				 "Sekarang berumur " + strconv.Itoa(umur) + " tahun"
}

func currentAge(tahun int, nama string) (currAge int, name string) {
	currAge = 2022 - tahun
	name = nama

	return
}

func main() {
	fmt.Println("function ini menghasilkan", multiply(2, 3));
	fmt.Println("function ini menghasilkan", multiply(2, 13));
	basicInfo, ageInfo := getBiography(22, "Didik Nur Hidayat", "Backend Developer")
	fmt.Println(basicInfo, ageInfo)
	age, name := currentAge(2000, "Didik")
	fmt.Println(name, age)
}

