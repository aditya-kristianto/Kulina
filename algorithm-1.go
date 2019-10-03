package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, _ := reader.ReadString('\n')
	n = strings.Replace(n, "\n", "", 1)
	nInteger, err := strconv.Atoi(n)
	if err != nil {
		fmt.Printf("%q isn't a integer.\n", n)

		return
	}

	ar, _ := reader.ReadString('\n')
	ar = strings.Replace(ar, "\n", "", 1)
	arArray := strings.Split(ar, " ")

	if nInteger != len(arArray) {
		fmt.Printf("Array count isn't matched.\n")

		return
	}

	arArrayInt := make([]int, nInteger)
	for i := range arArrayInt {
		arArrayInt[i], err = strconv.Atoi(arArray[i])
		if err != nil {
			fmt.Printf("%q isn't a integer.\n", arArray[i])

			return
		}
	}

	sort.Ints(arArrayInt)

	var tempInt int = 0
	var tempCount int = 0
	var total int = 0
	for i := range arArrayInt {
		if tempInt != arArrayInt[i] {
			total += tempCount / 2

			tempInt = arArrayInt[i]
			tempCount = 1
		} else {
			tempCount++
		}

		if i == len(arArrayInt)-1 {
			total += tempCount / 2
		}
	}

	fmt.Println(total)
}
