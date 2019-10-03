package main

import "fmt"

var animal string = "cat"

func main() {
	switch animal {
	case "cat":
		fmt.Println("meow")
	case "dog":
		fmt.Println("wooogh")
	case "mouse":
		fmt.Println("ciiit")
	}
}
