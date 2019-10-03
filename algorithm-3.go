package main

import (
	"fmt"
	"strings"
)

var number = "1345679"

func main() {
	var length = len(number)
	for i := 0; i < length; i++ {
		fmt.Printf("%c"+strings.Repeat("0", length-(i+1))+"\n", number[i])
	}
}
