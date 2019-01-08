package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var testdata = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var reader = bufio.NewReader(os.Stdin)

// use quickunion algorithm to find data, find cost is 1, union cost is N
func main() {
	for {
		fmt.Println("union or connected?")
		input, error := reader.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		} else {
			input = strings.TrimSpace(input)
			if strings.EqualFold("connected", input) {
				if isConnected(getInput()) {
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

// ~O(n) to find
func isConnected(x int, y int) bool {
	return findRoot(x) == findRoot(y)
}

// ~O(n) to union
func union(x int, y int) {
	xr := findRoot(x)
	fmt.Printf("before union, %d root is %d\n", x, xr)
	yr := findRoot(y)
	fmt.Printf("before union, %d root is %d\n", y, yr)
	testdata[xr] = testdata[yr]
	_xr := findRoot(x)
	_yr := findRoot(y)
	fmt.Printf("after union, %d root is %d\n", x, _xr)
	fmt.Printf("after union, %d root is %d\n", y, _yr)
}

func findRoot(x int) int {
	if x == testdata[x] {
		return x
	} else {
		return findRoot(testdata[x])
	}
}
