package main

import (
	"fmt"
)

func chengePoint(point *int) {
	*point = 300
}

func main() {
	point := 100
	var pointAddress *int = &point

	*pointAddress = 200
	chengePoint(&point)

	fmt.Println(pointAddress)
	fmt.Println(point)
	

}