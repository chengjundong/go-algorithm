package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var testdata = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var reader = bufio.NewReader(os.Stdin)

// use quickfind algorithm to find data, find cost is 1, union cost is N
func main() {
	for {
		fmt.Println("union or connected?")
		input, error := reader.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		} else {
			input = strings.TrimSpace(input)
			if strings.EqualFold("connected", input) {
				if find(getInput()) {
					fmt.Println(">>>>>>>>>>>>>>>> they are connected!")
				} else {
					fmt.Println(">>>>>>>>>>>>>>>> they are not connected yet!")
				}
			} else if strings.EqualFold("union", input) {
				union(getInput())
			} else {
				log.Fatal(">>>>>>>>>>>>>>>> unknown commend!")
			}
		}
	}
}

// get the input from console
func getInput() (int, int) {
	fmt.Println("first data?")
	first, _ := reader.ReadString('\n')
	x, _ := strconv.Atoi(strings.TrimSpace(first))
	fmt.Println("second data?")
	second, _ := reader.ReadString('\n')
	y, _ := strconv.Atoi(strings.TrimSpace(second))
	return x, y
}

// O(1) to find
func find(x int, y int) bool {
	return testdata[x] == testdata[y]
}

// O(n) to union
func union(x int, y int) {
	ref := testdata[x]
	for i := len(testdata) - 1; i >= 0; i-- {
		if testdata[i] == ref {
			testdata[i] = testdata[y]
		}
	}
	fmt.Println("array afte union:", testdata)

}
