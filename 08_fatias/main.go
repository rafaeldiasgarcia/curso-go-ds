package main

import "fmt"

func main() {
	var x [100]int // isso é um array

	fatia := x[:10]
	fmt.Printf("x = %T | fatia = %T\n", x, fatia)
	fmt.Println("Tamanho fatia:", len(fatia))

	y := []int{} // isso é uma fatia
	fmt.Printf("x = %T\n", y)
	fmt.Println("Tamanho y:", len(y))
}
