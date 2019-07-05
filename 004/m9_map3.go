package main

import "fmt"

func main() {

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name": "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name": "Helium",
			"state": "gas",
		},
		"C": map[string]string{
			"name": "Carbon",
			"state": "solid",
		},
	}


	if el, ok := elements["C"]; ok {
		fmt.Println(el["name"], el["state"])
	}else {
		fmt.Println("not found")
	}
}
