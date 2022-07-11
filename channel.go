package main

import (
	"fmt"
	"sync"
	"time"
)

var itemsChannel = make(chan string)
var cleanedItemsChannel = make(chan string)

func main() {

	
	items := [7]string { "harta","kerang","batu","harta","kerang","batu","harta" }

	go penyelam(items)
	go membersihkan()
	go menyimpan()

	chann := make(chan string)
	go runChannel("Hallo", chann)

	for msg := range chann {
		fmt.Println(msg)
	}

	time.Sleep(500 * time.Millisecond)
}

func printtext(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
}

func penyelam(items [7]string) {
	for _, item := range items {
		if item == "harta" {
			fmt.Println("berhasil mendapatkan " + item)
			itemsChannel <- item
		}
	}
}

func runChannel(text string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- text
	}

	close(c)
}

func membersihkan() {
	for rawItem := range itemsChannel {
		fmt.Println("berhasil membersihkan " + rawItem)
		cleanedItemsChannel <- "harta bersih"
	}
}

func menyimpan() {
	for cleanedItem := range cleanedItemsChannel {
		fmt.Println("Berhasil menyimpan " + cleanedItem)
	}
}