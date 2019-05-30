package main

// использование переменных

import (
	"fmt"
)

func main() {

	x := "Hello"
	y := "World"
	fmt.Println(x == y)

	x = "Hello"
	y = "Hello"
	fmt.Println(x == y)
}