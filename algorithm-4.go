package main

import (
	"fmt"
)

func main() {
	var number [100]bool

	doTrips(100, &number)

	for i := 0; i < len(number); i++ {
		fmt.Printf("%d %t\n", i+1, number[i])
	}
}

func doTrips(tripCount int, number *[100]bool) {
	for i := 0; i < tripCount; i++ {
		switchedOnAll(number)
		switcheOnMultiplyByTwo(number)
		switcheOnMultiplyByThree(number)
	}
}

func switchedOnAll(number *[100]bool) {
	for i := 0; i < len(number); i++ {
		number[i] = true
	}
}

func switcheOnMultiplyByTwo(number *[100]bool) {
	for i := 0; i < len(number); i++ {
		if (i+1) > 2 && (i+1)%2 == 0 {
			number[i] = !number[i]
		}
	}
}

func switcheOnMultiplyByThree(number *[100]bool) {
	for i := 0; i < len(number); i++ {
		if (i+1) > 3 && (i+1)%3 == 0 {
			number[i] = !number[i]
		}
	}
}
