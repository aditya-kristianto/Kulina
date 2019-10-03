package main

import (
	"bufio"
	"fmt"
	"os"
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

	s, _ := reader.ReadString('\n')
	s = strings.Replace(s, "\n", "", 1)
	sArray := strings.Split(s, " ")

	if nInteger != len(sArray) {
		fmt.Printf("Array count isn't matched.\n")

		return
	}

	var level int = 0
	var isCount bool = true
	var totalMount int = 0
	var totalValley int = 0

	for i := range sArray {
		if sArray[i] != "D" && sArray[i] != "U" {
			fmt.Printf("%q isn't a valid parameter (D or U).\n", sArray[i])

			return
		}

		if sArray[i] == "D" {
			level--
		} else if sArray[i] == "U" {
			level++
		}

		if level == 1 && isCount {
			isCount = false
			totalMount++
		} else if level == -1 && isCount {
			isCount = false
			totalValley++
		} else if level == 0 {
			isCount = true
		}
	}

	fmt.Println(totalValley)
}
