package main

import (
	"fmt"
	"strconv"
)

func bayarUtang(uang int) (pesan string) {
	utang := 50000

	switch {
		case uang > utang:
			pesan = "uang kamu kebanyakan"
		case uang == utang:
			pesan = "uangnya pass ya"
		default:
			pesan = "uang kamu kurang Rp." + strconv.Itoa(utang - uang)
	}

	return
}

func cekLampu(warna string) (pesan string) {
	switch warna {
	case "merah":
		pesan = "Berhenti !!"
	case "kuning":
		pesan = "Hati - Hati !"
	case "hijau":
		pesan = "Boleh Jalan"
	default:
		pesan = "Lampu error !!"
	}

	return
}

func main() {
	fmt.Println(bayarUtang(100000));
	fmt.Println(bayarUtang(50000));
	fmt.Println(bayarUtang(42500));

	// function cekLampu
	fmt.Println(cekLampu("abu-abu"))
}