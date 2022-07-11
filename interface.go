package main

import (
	"fmt"
)

type Calculator interface {
	penjumlahan() float32
	perkalian() float32
	pengurangan() float32
	pembagian() float32
}

type Value struct {
	num1 float32
	num2 float32
}

func (v Value) penjumlahan() float32 {
	return v.num1 + v.num2
}

func (v Value) perkalian() float32 {
	return v.num1 * v.num2
}

func (v Value) pengurangan() float32 {
	return v.num1 - v.num2
}

func (v Value) pembagian() float32 {
	return v.num1 / v.num2
}

func main() {
	val := Value{1,2}

	fmt.Println(val.penjumlahan())
	fmt.Println(val.perkalian())
	fmt.Println(val.pengurangan())
	fmt.Println(val.pembagian())
}