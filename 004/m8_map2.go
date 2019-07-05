package main

import "fmt"

func main() {
	//elements := make(map[string]string)
	//elements["H"] = "Hydrogen"
	//elements["He"] = "Helium"

	elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
	}



	fmt.Println(elements["H"])
	fmt.Println(elements["Un"])

	//name, ok := elements["Un"]
	//fmt.Println(name, ok)

	if name, ok := elements["un"]; ok {
		fmt.Println(name, ok)
	}else {
		fmt.Println("not found")
	}
}
