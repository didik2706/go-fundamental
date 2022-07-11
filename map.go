package main

import (
	"fmt"
)

var mahasiswa = map[int]string {
	1821400025: "Didik Nur Hidayat",
	1821400093: "Abdul Karim",
	1821400024: "Ahmad Baidawi",
}

func main() {
	// cek value map
	// name, exist := mahasiswa[1821400025]
	// fmt.Println(name, exist)
	if checkMahasiswa(1821400025) {
		fmt.Println("Mahasiswa dengan nim ini ada")
	} else {
		fmt.Println("Mahasiswa dengan nim ini tidak ada")
	}

	// menghapus value dari map
	delete(mahasiswa, 1821400024)

	// fmt.Println(mahasiswa[1821400125])
	// menampilkan map menggunakan for range
	for nim, nama := range mahasiswa {
		fmt.Println(nim, nama)
	}
}

func checkMahasiswa(nim int) bool {
	_, exist := mahasiswa[nim]
	
	return exist
}