package main

import "fmt"

func main() {
	//var x map[string]int
	// ошибка при исполнении
	// карта должна быть проинициализирована

	x := make(map[string]int)

	x["key"] = 10

	fmt.Println(x)
	fmt.Println(x["key"])
}
